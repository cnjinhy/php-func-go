# php-func-go

#### Description
Some `PHP` built-in functions implemented using Golang

使用golang实现的一些PHP函数，方便开发使用。

Note: Due to the language features of `Golang` and `PHP`, some functions are not 100% compatible. For example, the maps in Golang are unordered, and when using `http_ build_query`, the output order may be inconsistent. For example, arrays in Golang cannot be accessed outside the bounds, so there are restrictions on `empty`, and so on.

注意：由于`Golang`和`PHP`的语言特性，有些函数并不是100%兼容，比如Golang中的map是无序的，当使用`http_build_query`的时候有可能输出顺序不一致，比如Golang中的数组不能越界访问，所以empty也有所限制，等等。

#### Demo
the `README.md` is generated using the `php-func-go` function by [doc.go](https://github.com/cnjinhy/php-func-go/blob/master/doc.go)

`README.md`文件是通过`doc.go`生成的，`doc.go`使用了php-func-go的函数写的。

#### Unit Test
You can run `go test` in the dir `tests`

#### Function list
| php function | golang function | input argvs | return type |
|-------------|--------------|----------------------|--------|
| addslashes | Addslashes | s string | string  |
| array_key_exists | ArrayKeyExists | key interface{}, array interface{} | bool  |
| base64_decode | Base64Decode | encoded string | string  |
| base64_encode | Base64Encode | str string | string  |
| basename | Basename | path string, suffix ...string | string  |
| bin2hex | Bin2hex | input string | string  |
| checkdate | Checkdate | month, day, year int | bool  |
| chop | Chop | s string, charsToRemove ...string | string  |
| chr | Chr | codepoint int | string  |
| copy | Copy | src string, dest string | bool  |
| crc32 | Crc32 | s string | uint32  |
| date | Date | format string, timestamp ...interface{} | string  |
| dirname | Dirname | path string, levels ...int | string  |
| echo | Echo | args ...interface{} |  |
| empty | Empty | val interface{} | bool  |
| explode | Explode | delimiter string, str string | []string  |
| file_exists | FileExists | filename string | bool  |
| file_get_contents | FileGetContents | filename string | string  |
| file_put_contents | FilePutContents | filename string, data string, append ...int | bool  |
| filectime | Filectime | path string | int64  |
| filemtime | Filemtime | path string | int64  |
| filesize | Filesize | filename string | int64  |
| getrandmax | Getrandmax |  | int  |
| gettype | Gettype | any interface{} | string  |
| glob | Glob | pattern string | []string  |
| html_entity_decode | HtmlEntityDecode | s string | string  |
| htmlentities | Htmlentities | s string | string  |
| htmlspecialchars | Htmlspecialchars | s string | string  |
| htmlspecialchars_decode | HtmlspecialcharsDecode | s string | string  |
| http_build_query | HttpBuildQuery | data interface{} | string  |
| implode | Implode | glue string, pieces []string | string  |
| in_array | InArray | needle interface{}, haystack interface{}, strict ...bool | bool  |
| is_dir | IsDir | path string | bool  |
| is_numeric | IsNumeric | value interface{} | bool  |
| is_readable | IsReadable | filename string | bool  |
| json_encode | JsonEncode | value interface{} | string  |
| lcfirst | Lcfirst | str string | string  |
| lcg_value | LcgValue |  | float64  |
| ltrim | Ltrim | s string, charsToRemove ...string | string  |
| mb_strlen | MbStrlen | str string | int  |
| mb_substr | MbSubstr | str string, start int, length ...int | string  |
| md5 | Md5 | args ...interface{} | string  |
| md5_file | Md5File | path string, prevHash ...string | string  |
| microtime | Microtime | getAsFloat ...bool | interface  |
| mt_rand | MtRand | min int, max int | int  |
| nl2br | Nl2br | s string, useXHTML ...bool | string  |
| ord | Ord | s string | int  |
| parse_url | ParseUrl | urlStr string | *urlInfo  |
| preg_replace | PregReplace | pattern string, replacement string, subject string | string  |
| preg_replace_callback | PregReplaceCallback | pattern string, subject string, callback func | string |
| rawurldecode | Rawurldecode | str string | (string, error |
| rawurlencode | Rawurlencode | str string | string  |
| rtrim | Rtrim | s string, charsToRemove ...string | string  |
| sha1 | Sha1 | s string | string  |
| sha1_file | Sha1File | filename string, binary ...bool | string  |
| sleep | Sleep | seconds int |  |
| str_contains | StrContains | haystack string, needle string | bool  |
| str_ends_with | StrEndsWith | haystack string, needle string | bool  |
| str_repeat | StrRepeat | input string, time int | string  |
| str_replace | StrReplace | search interface{}, replace string, subject interface{} | string  |
| str_starts_with | StrStartsWith | haystack string, needle string | bool  |
| str_word_count | StrWordCount | s string, format ...int | interface  |
| strip_tags | StripTags | s string | string  |
| stripos | Stripos | haystack string, needle string | interface  |
| strlen | Strlen | str string | int  |
| strpos | Strpos | haystack string, needle string | interface  |
| strrev | Strrev | str string | string  |
| strstr | Strstr | s, substr string, options ...bool | string  |
| strtolower | Strtolower | s string | string  |
| strtotime | Strtotime | dateStr string | int64  |
| strtoupper | Strtoupper | s string | string  |
| strval | Strval | any interface{} | string  |
| substr | Substr | str string, start int, length ...int | (substr string |
| substr_count | SubstrCount | s, sub string, args ...int | int  |
| substr_replace | SubstrReplace | str, replace string, start int, length ...int | string  |
| time | Time |  | int64  |
| time_sleep_until | TimeSleepUntil | timestamp int64 |  |
| trim | Trim | str string, charlist ...string | string  |
| ucfirst | Ucfirst | s string | string  |
| ucwords | Ucwords | s string | string  |
| urldecode | Urldecode | str string | string  |
| urlencode | Urlencode | str string | string  |
| usleep | Usleep | microseconds int |  |
| var_dump | VarDump | v interface{} |  |
