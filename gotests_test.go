package gotests

import (
	"encoding/json"
	"errors"
	"flag"
	"go/types"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"unicode"

	"golang.org/x/tools/imports"

	"github.com/cweill/gotests"
)

var update = flag.Bool("update-golden-files", false, "Update golden files.")

func TestGenerateTests(t *testing.T) {
	type args struct {
		srcPath            string
		only               *regexp.Regexp
		excl               *regexp.Regexp
		exported           bool
		printInputs        bool
		subtests           bool
		parallel           bool
		named              bool
		importer           types.Importer
		template           string
		templateParamsPath string
		templateData       [][]byte
	}
	tests := []struct {
		name              string
		args              args
		wantFile          string
		wantNoTests       bool
		wantMultipleTests bool
		wantErr           bool
	}{
		{
			name: "Blank Go file",
			args: args{
				srcPath: `testdata/blankfile/blank.go`,
			},
			wantNoTests: true,
			wantErr:     true,
		},
		{
			name: "Blank Go file in directory",
			args: args{
				srcPath: `testdata/blankfile/notblank.go`,
			},
			wantNoTests: true,
			wantErr:     true,
		},
		{
			name: "Test file with garbage data",
			args: args{
				srcPath: `testdata/invalidtest/invalid.go`,
			},
			wantNoTests: true,
			wantErr:     true,
		},
		{
			name: "Hidden file",
			args: args{
				srcPath: `testdata/.hidden.go`,
			},
			wantNoTests: true,
			wantErr:     true,
		},
		{
			name: "Nonexistant file",
			args: args{
				srcPath: `testdata/nonexistant.go`,
			},
			wantNoTests: true,
			wantErr:     true,
		},
		{
			name: "Target test file",
			args: args{
				srcPath:  `testdata/test103_test.go`,
				only:     regexp.MustCompile("wrapToString"),
				subtests: true,
			},
			wantNoTests: false,
			wantErr:     false,
			wantFile:    `testdata/goldens/target_test_file.go`,
		},
		{
			name: "Target test file without only flag",
			args: args{
				srcPath:  `testdata/test103_test.go`,
				subtests: true,
			},
			wantNoTests: false,
			wantErr:     false,
			wantFile:    `testdata/goldens/target_test_file.go`,
		},
		{
			name: "No funcs",
			args: args{
				srcPath: `testdata/test000.go`,
			},
			wantNoTests: true,
		},
		{
			name: "Function with neither receiver, parameters, nor results",
			args: args{
				srcPath: `testdata/test001.go`,
			},
			wantFile: "testdata/goldens/function_with_neither_receiver_parameters_nor_results.go",
		},
		{
			name: "Function with anonymous arguments",
			args: args{
				srcPath: `testdata/test002.go`,
			},
			wantFile: "testdata/goldens/function_with_anonymous_arguments.go",
		},
		{
			name: "Function with named argument",
			args: args{
				srcPath: `testdata/test003.go`,
			},
			wantFile: "testdata/goldens/function_with_named_argument.go",
		},
		{
			name: "Function with return value",
			args: args{
				srcPath: `testdata/test004.go`,
			},
			wantFile: "testdata/goldens/function_with_return_value.go",
		},
		{
			name: "Function returning an error",
			args: args{
				srcPath: `testdata/test005.go`,
			},
			wantFile: "testdata/goldens/function_returning_an_error.go",
		},
		{
			name: "Function with multiple arguments",
			args: args{
				srcPath: `testdata/test006.go`,
			},
			wantFile: "testdata/goldens/function_with_multiple_arguments.go",
		},
		{
			name: "Print inputs with multiple arguments",
			args: args{
				srcPath:     `testdata/test006.go`,
				printInputs: true,
			},
			wantFile: "testdata/goldens/print_inputs_with_multiple_arguments.go",
		},
		{
			name: "Method on a struct pointer",
			args: args{
				srcPath: `testdata/test007.go`,
			},
			wantFile: "testdata/goldens/method_on_a_struct_pointer.go",
		},
		{
			name: "Print inputs with single argument",
			args: args{
				srcPath:     `testdata/test007.go`,
				printInputs: true,
			},
			wantFile: "testdata/goldens/print_inputs_with_single_argument.go",
		},
		{
			name: "Function with struct pointer argument and return type",
			args: args{
				srcPath: `testdata/test008.go`,
			},
			wantFile: "testdata/goldens/function_with_struct_pointer_argument_and_return_type.go",
		},
		{
			name: "Struct value method with struct value return type",
			args: args{
				srcPath: `testdata/test009.go`,
			},
			wantFile: "testdata/goldens/struct_value_method_with_struct_value_return_type.go",
		},
		{
			name: "Function with map argument and return type",
			args: args{
				srcPath: `testdata/test010.go`,
			},
			wantFile: "testdata/goldens/function_with_map_argument_and_return_type.go",
		},
		{
			name: "Function with slice argument and return type",
			args: args{
				srcPath: `testdata/test011.go`,
			},
			wantFile: "testdata/goldens/function_with_slice_argument_and_return_type.go",
		},
		{
			name: "Function returning only an error",
			args: args{
				srcPath: `testdata/test012.go`,
			},
			wantFile: "testdata/goldens/function_returning_only_an_error.go",
		},
		{
			name: "Function with a function parameter",
			args: args{
				srcPath: `testdata/test013.go`,
			},
			wantFile: "testdata/goldens/function_with_a_function_parameter.go",
		},
		{
			name: "Function with a function parameter with its own parameters and result",
			args: args{
				srcPath: `testdata/test014.go`,
			},
			wantFile: "testdata/goldens/function_with_a_function_parameter_with_its_own_parameters_and_result.go",
		},
		{
			name: "Function with a function parameter that returns two results",
			args: args{
				srcPath: `testdata/test015.go`,
			},
			wantFile: "testdata/goldens/function_with_a_function_parameter_that_returns_two_results.go",
		},
		{
			name: "Function with defined interface type parameter and result",
			args: args{
				srcPath: `testdata/test016.go`,
			},
			wantFile: "testdata/goldens/function_with_defined_interface_type_parameter_and_result.go",
		},
		{
			name: "Function with imported interface receiver, parameter, and result",
			args: args{
				srcPath: `testdata/test017.go`,
			},
			wantFile: "testdata/goldens/function_with_imported_interface_receiver_parameter_and_result.go",
		},
		{
			name: "Function with imported struct receiver, parameter, and result",
			args: args{
				srcPath: `testdata/test018.go`,
			},
			wantFile: "testdata/goldens/function_with_imported_struct_receiver_parameter_and_result.go",
		},
		{
			name: "Function with multiple parameters of the same type",
			args: args{
				srcPath: `testdata/test019.go`,
			},
			wantFile: "testdata/goldens/function_with_multiple_parameters_of_the_same_type.go",
		},
		{
			name: "Function with a variadic parameter",
			args: args{
				srcPath: `testdata/test020.go`,
			},
			wantFile: "testdata/goldens/function_with_a_variadic_parameter.go",
		},
		{
			name: "Function with interface{} parameter and result",
			args: args{
				srcPath: `testdata/test021.go`,
			},
			wantFile: "testdata/goldens/function_with_interface_parameter_and_result.go",
		},
		{
			name: "Function with named imports",
			args: args{
				srcPath: `testdata/test022.go`,
			},
			wantFile: "testdata/goldens/function_with_named_imports.go",
		},
		{
			name: "Function with channel parameter and result",
			args: args{
				srcPath: `testdata/test023.go`,
			},
			wantFile: "testdata/goldens/function_with_channel_parameter_and_result.go",
		},
		{
			name: "File with multiple imports",
			args: args{
				srcPath: `testdata/test024.go`,
			},
			wantFile: "testdata/goldens/file_with_multiple_imports.go",
		},
		{
			name: "Function returning two results and an error",
			args: args{
				srcPath: `testdata/test025.go`,
			},
			wantFile: "testdata/goldens/function_returning_two_results_and_an_error.go",
		},
		{
			name: "Multiple named results",
			args: args{
				srcPath: `testdata/test026.go`,
			},
			wantFile: "testdata/goldens/multiple_named_results.go",
		},
		{
			name: "Two different structs with same method name",
			args: args{
				srcPath: `testdata/test027.go`,
			},
			wantFile: "testdata/goldens/two_different_structs_with_same_method_name.go",
		},
		{
			name: "Underlying types",
			args: args{
				srcPath: `testdata/test028.go`,
			},
			wantFile: "testdata/goldens/underlying_types.go",
		},
		{
			name: "Struct receiver with multiple fields",
			args: args{
				srcPath: `testdata/test029.go`,
			},
			wantFile: "testdata/goldens/struct_receiver_with_multiple_fields.go",
		},
		{
			name: "Struct receiver with anonymous fields",
			args: args{
				srcPath: `testdata/test030.go`,
			},
			wantFile: "testdata/goldens/struct_receiver_with_anonymous_fields.go",
		},
		{
			name: "io.Writer parameters",
			args: args{
				srcPath: `testdata/test031.go`,
			},
			wantFile: "testdata/goldens/io_writer_parameters.go",
		},
		{
			name: "Two structs with same method name",
			args: args{
				srcPath: `testdata/test032.go`,
			},
			wantFile: "testdata/goldens/two_structs_with_same_method_name.go",
		},
		{
			name: "Functions and methods with 'name' receivers, parameters, and results",
			args: args{
				srcPath: `testdata/test033.go`,
			},
			wantFile: "testdata/goldens/functions_and_methods_with_name_receivers_parameters_and_results.go",
		},
		{
			name: "Receiver struct with reserved field names",
			args: args{
				srcPath: `testdata/test034.go`,
			},
			wantFile: "testdata/goldens/receiver_struct_with_reserved_field_names.go",
		},
		{
			name: "Receiver struct with fields with complex package names",
			args: args{
				srcPath: `testdata/test035.go`,
			},
			wantFile: "testdata/goldens/receiver_struct_with_fields_with_complex_package_names.go",
		},
		{
			name: "Functions and receivers with same names except exporting",
			args: args{
				srcPath: `testdata/test036.go`,
			},
			wantFile: "testdata/goldens/functions_and_receivers_with_same_names_except_exporting.go",
		},
		{
			name: "Receiver is indirect imported struct",
			args: args{
				// only:    regexp.MustCompile("^Foo037$"),
				srcPath: `testdata/test037.go`,
			},
			wantFile: "testdata/goldens/receiver_is_indirect_imported_struct.go",
		},
		{
			name: "Multiple functions",
			args: args{
				srcPath: `testdata/test_filter.go`,
			},
			wantFile: "testdata/goldens/multiple_functions.go",
		},
		{
			name: "Multiple functions with only",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("FooFilter|bazFilter"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_only.go",
		},
		{
			name: "Multiple functions with only regexp without matches",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("asdf"),
			},
			wantNoTests: true,
		},
		{
			name: "Multiple functions with case-insensitive only",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("(?i)fooFilter|BazFilter"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_case-insensitive_only.go",
		},
		{
			name: "Multiple functions with only filtering on receiver",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("^BarBarFilter$"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_only_filtering_on_receiver.go",
		},
		{
			name: "Multiple functions with only filtering on method",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("^(BarFilter)$"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_only_filtering_on_method.go",
		},
		{
			name: "Multiple functions filtering exported",
			args: args{
				srcPath:  `testdata/test_filter.go`,
				exported: true,
			},
			wantFile: "testdata/goldens/multiple_functions_filtering_exported.go",
		},
		{
			name: "Multiple functions filtering exported with only",
			args: args{
				srcPath:  `testdata/test_filter.go`,
				only:     regexp.MustCompile(`FooFilter`),
				exported: true,
			},
			wantFile: "testdata/goldens/multiple_functions_filtering_exported_with_only.go",
		},
		{
			name: "Multiple functions filtering all out",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("fooFilter"),
			},
			wantNoTests: true,
		},
		{
			name: "Multiple functions with excl",
			args: args{
				srcPath: `testdata/test_filter.go`,
				excl:    regexp.MustCompile("FooFilter|bazFilter"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_excl.go",
		},
		{
			name: "Multiple functions with case-insensitive excl",
			args: args{
				srcPath: `testdata/test_filter.go`,
				excl:    regexp.MustCompile("(?i)foOFilter|BaZFilter"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_case-insensitive_excl.go",
		},
		{
			name: "Multiple functions filtering exported with excl",
			args: args{
				srcPath:  `testdata/test_filter.go`,
				excl:     regexp.MustCompile(`FooFilter`),
				exported: true,
			},
			wantFile: "testdata/goldens/multiple_functions_filtering_exported_with_excl.go",
		},
		{
			name: "Multiple functions excluding all",
			args: args{
				srcPath: `testdata/test_filter.go`,
				excl:    regexp.MustCompile("bazFilter|FooFilter|BarFilter"),
			},
			wantNoTests: true,
		},
		{
			name: "Multiple functions excluding on receiver",
			args: args{
				srcPath: `testdata/test_filter.go`,
				excl:    regexp.MustCompile("^BarBarFilter$"),
			},
			wantFile: "testdata/goldens/multiple_functions_excluding_on_receiver.go",
		},
		{
			name: "Multiple functions excluding on method",
			args: args{
				srcPath: `testdata/test_filter.go`,
				excl:    regexp.MustCompile("^BarFilter$"),
			},
			wantFile: "testdata/goldens/multiple_functions_excluding_on_method.go",
		},
		{
			name: "Multiple functions with both only and excl",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("BarFilter"),
				excl:    regexp.MustCompile("FooFilter"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_both_only_and_excl.go",
		},
		{
			name: "Multiple functions with only and excl competing",
			args: args{
				srcPath: `testdata/test_filter.go`,
				only:    regexp.MustCompile("FooFilter|BarFilter"),
				excl:    regexp.MustCompile("FooFilter|bazFilter"),
			},
			wantFile: "testdata/goldens/multiple_functions_with_only_and_excl_competing.go",
		},
		{
			name: "Custom importer fails",
			args: args{
				srcPath:  `testdata/test_filter.go`,
				importer: &fakeImporter{err: errors.New("error")},
			},
			wantFile: "testdata/goldens/custom_importer_fails.go",
		},
		{
			name: "Existing test file",
			args: args{
				srcPath: `testdata/test100.go`,
			},
			wantFile: "testdata/goldens/existing_test_file.go",
		},
		{
			name: "Existing test file with just package declaration",
			args: args{
				srcPath: `testdata/test101.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_just_package_declaration.go",
		},
		{
			name: "Existing test file with no functions",
			args: args{
				srcPath: `testdata/test102.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_no_functions.go",
		},
		{
			name: "Existing test file with multiple imports",
			args: args{
				srcPath: `testdata/test200.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_multiple_imports.go",
		},
		{
			name: "Different packages in same directory - part 1",
			args: args{
				srcPath: `testdata/mixedpkg/bar.go`,
			},
			wantFile: "testdata/goldens/different_packages_in_same_directory_-_part_1.go",
		},
		{
			name: "Different packages in same directory - part 2",
			args: args{
				srcPath: `testdata/mixedpkg/foo.go`,
			},
			wantFile: "testdata/goldens/different_packages_in_same_directory_-_part_2.go",
		},
		{
			name: "Empty test file",
			args: args{
				srcPath: `testdata/blanktest/blank.go`,
			},
			wantFile: "testdata/goldens/empty_test_file.go",
		},
		{
			name: "Test file with syntax errors",
			args: args{
				srcPath: `testdata/syntaxtest/syntax.go`,
			},
			wantFile: "testdata/goldens/test_file_with_syntax_errors.go",
		},
		{
			name: "Undefined types",
			args: args{
				srcPath: `testdata/undefinedtypes/undefined.go`,
			},
			wantFile: "testdata/goldens/undefined_types.go",
		},
		{
			name: "Subtest Edition - Functions and receivers with same names except exporting",
			args: args{
				srcPath:  `testdata/test036.go`,
				subtests: true,
			},
			wantFile: "testdata/goldens/subtest_edition_-_functions_and_receivers_with_same_names_except_exporting.go",
		},
		{
			name: "Test t method receiver",
			args: args{
				srcPath: `testdata/test_t_receiver.go`,
			},
			wantNoTests: false,
			wantErr:     false,
			wantFile:    `testdata/goldens/test_t_receiver.go`,
		},
		{
			name: "Init function",
			args: args{
				srcPath: `testdata/init_func.go`,
			},
			wantFile: "testdata/goldens/no_init_funcs.go",
		},
		{
			name: "Existing test file with package level comments",
			args: args{
				srcPath: `testdata/test_existing_test_file_with_comments.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_package_level_comments.go",
		},
		{
			name: "Existing test file with package level comments with newline",
			args: args{
				srcPath: `testdata/test_existing_test_file_with_comments.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_package_level_comments.go",
		},
		{
			name: "Existing test file with package level comments without newline",
			args: args{
				srcPath: `testdata/test_existing_test_file_with_comments_without_newline.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_package_level_comments_without_newline.go",
		},
		{
			name: "Existing test file with mixed types package level comments",
			args: args{
				srcPath: `testdata/test_existing_test_file_with_mixed_comments.go`,
			},
			wantFile: "testdata/goldens/existing_test_file_with_package_level_mixed_types_comments.go",
		},
		{
			name: "Naked function with subtests",
			args: args{
				srcPath:  "testdata/naked_function_with_subtests.go",
				subtests: true,
			},
			wantFile: "testdata/goldens/naked_function_with_subtests.go",
		},
		{
			name: "Naked function with parallel subtests",
			args: args{
				srcPath:  "testdata/naked_function_with_parallel_subtests.go",
				subtests: true,
				parallel: true,
			},
			wantFile: "testdata/goldens/naked_function_with_parallel_subtests.go",
		},
		{
			name: "Naked function without subtests",
			args: args{
				srcPath:  "testdata/naked_function_without_subtests.go",
				subtests: false,
			},
			wantFile: "testdata/goldens/naked_function_without_subtests.go",
		},
		{
			name: "Naked function without subtests with parallel",
			args: args{
				srcPath:  "testdata/naked_function_without_subtests_with_parallel.go",
				parallel: true,
			},
			wantFile: "testdata/goldens/naked_function_without_subtests_with_parallel.go",
		},
		{
			name: "Test interface embedding",
			args: args{
				srcPath: `testdata/undefinedtypes/interface_embedding.go`,
			},
			wantFile:    "testdata/goldens/interface_embedding.go",
			wantNoTests: !versionGreaterOrEqualThan("go1.11"),
			wantErr:     !versionGreaterOrEqualThan("go1.11"),
		},
		{
			name: "named_named=on",
			args: args{
				srcPath: "testdata/test038.go",
				named:   true,
			},
			wantFile: "testdata/named/named_on.go",
		},
		{
			name: "named_named=on,subtests=on",
			args: args{
				srcPath:  "testdata/test038.go",
				subtests: true,
				named:    true,
			},
			wantFile: "testdata/named/named_on_subtests_on.go",
		},
		{
			name: "named_named=on,subtests=on,parallel=on",
			args: args{
				srcPath:  "testdata/test038.go",
				subtests: true,
				parallel: true,
				named:    true,
			},
			wantFile: "testdata/named/named_on_subtests_on_parallel_on.go",
		},
	}
	tmp, err := ioutil.TempDir("", "gotests_test")
	if err != nil {
		t.Fatalf("ioutil.TempDir: %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var params map[string]interface{}
			var err error
			var want string
			if tt.args.templateParamsPath != "" {
				params, err = loadExternalJsonFile(tt.args.templateParamsPath)
				if err != nil {
					t.Error(tt.name, err)
					return
				}
			}
			gts, err := gotests.GenerateTests(tt.args.srcPath, &gotests.Options{
				Only:           tt.args.only,
				Exclude:        tt.args.excl,
				Exported:       tt.args.exported,
				PrintInputs:    tt.args.printInputs,
				Subtests:       tt.args.subtests,
				Parallel:       tt.args.parallel,
				Named:          tt.args.named,
				Importer:       func() types.Importer { return tt.args.importer },
				TemplateDir:    "cmp-linted",
				Template:       tt.args.template,
				TemplateParams: params,
				TemplateData:   tt.args.templateData,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("%q. GenerateTests(%v) error = %v, wantErr %v", tt.name, tt.args.srcPath, err, tt.wantErr)
				return
			}
			if (len(gts) == 0) != tt.wantNoTests {
				t.Errorf("%q. GenerateTests(%v) returned no tests", tt.name, tt.args.srcPath)
				return
			}
			if (len(gts) > 1) != tt.wantMultipleTests {
				t.Errorf("%q. GenerateTests(%v) returned too many tests", tt.name, tt.args.srcPath)
				return
			}
			if tt.wantNoTests || tt.wantMultipleTests {
				return
			}

			got := string(gts[0].Output)
			if *update {
				if err := ioutil.WriteFile(tt.wantFile, []byte(got), 0o644); err != nil {
					t.Fatalf("ioutil.WriteFile %s: %v", tt.wantFile, err)
				}
			}
			if tt.wantFile != "" {
				want = mustReadAndFormatGoFile(t, tt.wantFile)
			}
			if got != want {
				t.Errorf("%q. GenerateTests(%v) = \n%v, want \n%v", tt.name, tt.args.srcPath, got, want)
				outputResult(t, tmp, tt.name, gts[0].Output)
			}
		})
	}
}

func versionGreaterOrEqualThan(version string) bool {
	prefixes := []string{"go1.9", "go1.10", "go1.11", "go1.12", "go1.13"}
	v := runtime.Version()
	for _, prefix := range prefixes {
		if strings.Contains(version, prefix) {
			return true
		}
		if strings.Contains(v, prefix) {
			return false
		}
	}
	return true
}

func mustReadAndFormatGoFile(t *testing.T, filename string) string {
	fmted, err := imports.Process(filename, nil, nil)
	if err != nil {
		t.Fatalf("reading and formatting file: %v", err)
	}
	return string(fmted)
}

func outputResult(t *testing.T, tmpDir, testName string, got []byte) {
	tmpResult := path.Join(tmpDir, toSnakeCase(testName)+".go")
	if err := ioutil.WriteFile(tmpResult, got, 0o644); err != nil {
		t.Errorf("ioutil.WriteFile: %v", err)
	}
	t.Logf(tmpResult)
}

func loadExternalJsonFile(file string) (map[string]interface{}, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{}
	err = json.Unmarshal(buf, &params)
	return params, err
}

func mustLoadExternalTemplateDir(t *testing.T, dir string) [][]byte {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("ioutil.ReadDir: %v", err)
	}

	templateData := make([][]byte, 0)

	for _, f := range files {
		filePath := filepath.Join(dir, f.Name())
		templateData = append(templateData, mustLoadExternalTemplateFile(t, filePath))
	}

	return templateData
}

func mustLoadExternalTemplateFile(t *testing.T, file string) []byte {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatalf("loading external template file: %v", err)
	}

	return buf
}

func toSnakeCase(s string) string {
	var res []rune
	for _, r := range []rune(s) {
		r = unicode.ToLower(r)
		switch r {
		case ' ', '.':
			r = '_'
		case ',', '\'', '{', '}':
			continue
		}
		res = append(res, r)
	}
	return string(res)
}

// 249032394 ns/op
func BenchmarkGenerateTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.GenerateTests("testdata/", &gotests.Options{})
	}
}

// A fake importer.
type fakeImporter struct {
	err error
}

func (f *fakeImporter) Import(path string) (*types.Package, error) {
	return nil, f.err
}
