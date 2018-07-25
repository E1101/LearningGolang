package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Processes struct {
	Total    int      `ini:"total"`
	Running  int      `ini:"running"`
	Sleeping int      `ini:"sleeping"`
	Threads  int      `ini:"threads"`
	Load     float64  `ini:"load"`
}

func main() {
	fmt.Println("Write a struct to output:")

	// Marshals the struct
	//
	proc := &Processes{
		Total:    23,
		Running:  3,
		Sleeping: 20,
		Threads:  34,
		Load:     1.8,
	}

	data, err := Marshal(proc)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	// Creates a new Processes struct
	// and unmarshals the data into it
	//
	fmt.Println("Read the data back into a struct")
	proc2 := &Processes{}
	if err := Unmarshal(data, proc2); err != nil {
		panic(err)
	}

	// Prints out the struct
	fmt.Printf("Struct: %#v", proc2)
}


// Methods:

func Marshal(v interface{}) ([]byte, error) {
	var b bytes.Buffer

	// Gets a reflect.Value of the
	// current interface.
	// Dereferences pointers.
	val := reflect.Indirect(reflect.ValueOf(v))
	// handle only structs.
	if val.Kind() != reflect.Struct {
		return []byte{}, errors.New("unmarshal can only take structs")
	}

	// Loops through all of the
	// fields on the struct
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// Gets the name from tagName
		name := fieldName(f)
		// Relies on the print
		// formatter to print the
		// raw data into the buffer
		raw := val.Field(i).Interface()
		fmt.Fprintf(&b, "%s=%v\n", name, raw)
	}

	// Returns the contents of the buffer
	return b.Bytes(), nil
}

func Unmarshal(data []byte, v interface{}) error {
	val := reflect.Indirect(reflect.ValueOf(v))
	t := val.Type()
	b := bytes.NewBuffer(data)
	// From data, you use a scanner to
	// read one line of INI data at a time.
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		// Splits a line at the
		// equals sign
		line := scanner.Text()
		pair := strings.SplitN(line, "=", 2)
		if len(pair) < 2 {
			// Skip any malformed lines.
			continue
		}

		// Passes the task of setting
		// the value to setField()
		setField(pair[0], pair[1], t, val)
	}

	return nil
}


// protected

func setField(name, value string, t reflect.Type, v reflect.Value) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if name == fieldName(field) {
			var dest reflect.Value
			// Uses a kind switch to figure out
			// how to take your value string
			// and convert it to the right type
			switch field.Type.Kind() {

			// If you don’t know about the kind,
			// just skip the field. This isn’t an error
			default:
				fmt.Printf("Kind %s not supported.\n",
					field.Type.Kind())
				continue

			// This version supports only a few
			// kinds of values.
			// Supporting other types is usually easy,
			// but highly repetitive.

			// Once a raw value is converted to
			// its type, wraps it in a value
			case reflect.Int:
				ival, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf(
						"Could not convert %q to int: %s\n",
						value, err)
					continue
				}
				dest = reflect.ValueOf(ival)
			case reflect.Float64:
				fval, err := strconv.ParseFloat(value, 64)
				if err != nil {
					fmt.Printf(
						"Could not convert %q to float64: %s\n",
						value, err)
					continue
				}
				dest = reflect.ValueOf(fval)
			case reflect.String:
				dest = reflect.ValueOf(value)
			case reflect.Bool:
				bval, err := strconv.ParseBool(value)
				if err != nil {
					fmt.Printf(
						"Could not convert %q to bool: %s\n",
						value, err)
					continue
				}
				dest = reflect.ValueOf(bval)
			}

			// Sets the value for the
			// relevant struct field
			v.Field(i).Set(dest)
		}
	}
}

func fieldName(field reflect.StructField) string {
	// Gets the tag off the
	// struct field
	if t := field.Tag.Get("ini"); t != "" {
		return t
	}

	// If there is no tag, falls
	// back to the field name
	return field.Name
}
