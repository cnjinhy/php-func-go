#### php-func-go
Some PHP built-in functions implemented using Golang
#### desc
README.md is generated using the php-func-go function through doc.go
#### function list
| php function | golang function | input argvs | return type |
|-------------|--------------|----------------------|--------|
| addslashes | Addslashes | s string | string { |
| array_key_exists | ArrayKeyExists | key interface{}, array interface{} | bool { |
| base64_decode | Base64Decode | encoded string | string { |
| base64_encode | Base64Encode | str string | string { |
| basename | Basename | path string, suffix ...string | string { |
| bin2hex | Bin2hex | input string | string { |
| chop | Chop | s string, charsToRemove ...string | string { |
| chr | Chr | codepoint int | string { |
| copy | Copy | src string, dest string | bool { |
| crc32 | Crc32 | s string | uint32 { |
| date | Date | format string, timestamp ...interface{} | string { |
| dirname | Dirname | path string, levels ...int | string { |
| echo | Echo | args ...interface{} | { |
| explode | Explode | delimiter string, str string | []string { |
| file_exists | FileExists | filename string | bool { |
| file_get_contents | FileGetContents | filename string | string { |
| file_put_contents | FilePutContents | filename string, data string, append ...int | bool { |
| filesize | Filesize | filename string | (int64, error |
| getrandmax | Getrandmax |  | int { |
| glob | Glob | pattern string | []string { |
| http_build_query | HttpBuildQuery | data interface{} | string { |
| implode | Implode | glue string, pieces []string | string { |
| in_array | InArray | needle interface{}, haystack interface{}, strict ...bool | bool { |
| is_dir | IsDir | path string | bool { |
| is_numeric | IsNumeric | value interface{} | bool { |
| is_readable | IsReadable | filename string | bool { |
| json_encode | JsonEncode | value interface{} | string { |
| lcfirst | Lcfirst | str string | string { |
| lcg_value | LcgValue |  | float64 { |
| ltrim | Ltrim | s string, charsToRemove ...string | string { |
| mb_strlen | MbStrlen | str string | (strlen int |
| md5 | Md5 | args ...interface{} | string { |
| md5_file | Md5File | path string, prevHash ...string | string { |
| microtime | Microtime | getAsFloat ...bool | interface{} { |
| mt_rand | MtRand | min int, max int | int { |
| ord | Ord | s string | int { |
| parse_url | ParseUrl | urlStr string | *parseURLInfo { |
| preg_replace | PregReplace | pattern string, replacement string, subject string | string { |
| preg_replace_callback | PregReplaceCallback | pattern string, subject string, callback func | string |
| print_r | PrintR | vList ...interface{} | { |
| rtrim | Rtrim | s string, charsToRemove ...string | string { |
| sha1 | Sha1 | s string | string { |
| sha1_file | Sha1File | filename string, binary ...bool | string { |
| sleep | Sleep | seconds int | { |
| str_contains | StrContains | haystack string, needle string | bool { |
| str_ends_with | StrEndsWith | haystack string, needle string | bool { |
| str_replace | StrReplace | search, replace, subject interface{} | interface{} { |
| str_starts_with | StrStartsWith | haystack string, needle string | bool { |
| str_word_count | StrWordCount | s string, format ...int | interface{} { |
| strip_tags | StripTags | s string | string { |
| stripos | Stripos | haystack string, needle string | interface{} { |
| strlen | Strlen | str string | (strlen int |
| strpos | Strpos | haystack string, needle string | interface{} { |
| strrev | Strrev | str string | string { |
| strstr | Strstr | s, substr string, options ...bool | string { |
| strval | Strval | any interface{} | string { |
| substr_count | SubstrCount | s, sub string, args ...int | int { |
| substr_replace | SubstrReplace | str, replace string, start int, length ...int | string { |
| time | Time |  | int64 { |
| time_sleep_until | TimeSleepUntil | timestamp int64 | { |
| trim | Trim | str string, charlist ...string | string { |
| ucfirst | Ucfirst | s string | string { |
| ucwords | Ucwords | s string | string { |
| usleep | Usleep | microseconds int | { |
| var_dump | VarDump | v interface{} | { |
