package util

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

// set struct fields to flagset by filed's tag
func SetFlagsByStruct(flagSet *flag.FlagSet, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	if isStructPtr(objT) {
		return fmt.Errorf("%v must be a struct not a struct pointer!", obj)
	}
	return parseFlagsFromStruct("", flagSet, objT)
}

// set struct values from flagset
func SetValuesFromFlags(flagSet *flag.FlagSet, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	if !isStructPtr(objT) {
		return fmt.Errorf("%v must be  a struct pointer", obj)
	}
	objT = objT.Elem()
	objV = objV.Elem()

	return setValuesFromFlags("", flagSet, objT, objV)
}

func parseFlagsFromStruct(prefix string, flagSet *flag.FlagSet, objT reflect.Type) error {
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)

		opts := getFlagOptions(prefix, fieldT.Tag, fieldT.Type)

		if !fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			err := parseFlagsFromStruct(opts.Flag, flagSet, fieldT.Type)
			if err != nil {
				return err
			}
			continue
		}

		// set
		if err := opts.setToFlagSet(flagSet); err != nil {
			return err
		}
	}
	return nil
}

type flagOptions struct {
	Flag      string
	Shorthand string
	Type      reflect.Type
	Default   string
	Desc      string
}

func (opts flagOptions) setToFlagSet(flagSet *flag.FlagSet) error {
	flag := opts.Flag
	if flag == "" {
		return fmt.Errorf("flag should be a not empty string")
	}
	shorthand := opts.Shorthand
	def := defFromEnv(opts.Default)
	desc := opts.Desc

	switch opts.Type.Kind() {
	case reflect.String:
		flagSet.StringP(flag, shorthand, def, desc)
	case reflect.Bool:
		v := false
		lower := strings.ToLower(def)
		if lower == "yes" || lower == "1" || lower == "true" || lower == "on" {
			v = true
		}
		flagSet.BoolP(flag, shorthand, v, desc)
	case reflect.Int:
		v, err := strconv.Atoi(def)
		if err != nil {
			return fmt.Errorf("cannot convert value(%s) to int", def)
		}
		flagSet.IntP(flag, shorthand, v, desc)
	default:
		return fmt.Errorf("unsupport type(%s) to set flag(%s)", opts.Type.String(), flag)
	}
	return nil
}

func defFromEnv(def string) string {
	matchedParenthesis, err := regexp.MatchString(`^\$\(\w+\)$`, def)
	if err != nil {
		return def
	}

	matchedBrace, err := regexp.MatchString(`^\$\{\w+\}$`, def)
	if err != nil {
		return def
	}

	if matchedParenthesis || matchedBrace {
		return os.Getenv(def[2 : len(def)-1])
	}

	return def
}

func (opts flagOptions) setValuesFromFlagSet(fieldV reflect.Value, flagSet *flag.FlagSet) error {
	if opts.Flag == "" {
		return fmt.Errorf("flag should not be an empty string")
	}

	flag := opts.Flag
	switch opts.Type.Kind() {
	case reflect.String:
		val, err := flagSet.GetString(flag)
		if err != nil {
			return err
		}
		fieldV.Set(reflect.ValueOf(val))
	case reflect.Bool:
		val, err := flagSet.GetBool(flag)
		if err != nil {
			return err
		}
		fieldV.Set(reflect.ValueOf(val))
	case reflect.Int:
		val, err := flagSet.GetInt(flag)
		if err != nil {
			return err
		}
		fieldV.Set(reflect.ValueOf(val))
	default:
		return fmt.Errorf("unsupport type(%s) to set values form flag(%s)", opts.Type.String(), flag)
	}
	return nil
}

func getFlagOptions(prefix string, tag reflect.StructTag, typ reflect.Type) flagOptions {
	options := flagOptions{
		Type: typ,
	}
	// TODO optimize performance
	if s := tag.Get("flag"); s != "" {
		options.Flag = s
		if prefix != "" {
			options.Flag = prefix + "." + options.Flag
		}
	}
	if s := tag.Get("shorthand"); prefix == "" && s != "" {
		options.Shorthand = s
	}
	if s := tag.Get("default"); s != "" {
		options.Default = s
	}
	if s := tag.Get("desc"); s != "" {
		options.Desc = s
	}

	return options
}

func setValuesFromFlags(prefix string, flagSet *flag.FlagSet, objT reflect.Type, objV reflect.Value) error {

	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}
		fieldT := objT.Field(i)

		opts := getFlagOptions(prefix, fieldT.Tag, fieldT.Type)

		if !fieldT.Anonymous && fieldT.Type.Kind() == reflect.Struct {
			err := setValuesFromFlags(opts.Flag, flagSet, fieldT.Type, fieldV)
			if err != nil {
				return err
			}
			continue
		}

		if err := opts.setValuesFromFlagSet(fieldV, flagSet); err != nil {
			return err
		}

	}
	return nil
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
