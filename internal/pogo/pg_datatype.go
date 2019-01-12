package pogo

// DataType for a column
type DataType interface {
	String() string
}

// Postgres Numerics
// https://www.postgresql.org/docs/11/datatype-numeric.html

// SmallInt 2 bytes small-range integer -32768 to +32767
type SmallInt struct {
}

// Integer 4 bytes typical choice for integer -2147483648 to +2147483647
type Integer struct {
}

// BigInt 8 bytes large-range integer -9223372036854775808 to +9223372036854775807
type BigInt struct {
}

// Decimal variable user-specified precision, exact up to 131072 digits before the decimal point; up to 16383 digits after the decimal point
type Decimal struct {
}

// Numeric variable user-specified precision, exact up to 131072 digits before the decimal point; up to 16383 digits after the decimal point
type Numeric struct {
	// The precision must be positive, the scale zero or positive. Alternatively:
	// without any precision or scale creates a column in which numeric values of any precision and scale can be stored, up to the implementation limit on precision
	Precision int
	Scale     int
}

// Float is the SQL-standard notations float and float(p) for
// specifying inexact numeric types. Here, p specifies the minimum acceptable
// precision in binary digits. PostgreSQL accepts float(1) to float(24) as
// selecting the real type, while float(25) to float(53) select double precision.
// Values of p outside the allowed range draw an error. float with no precision
// specified is taken to mean double precision.
//
// Also accepted:
//
// Infinity
// -Infinity
// NaN
//
// When writing these values as constants in an SQL command, you must put quotes
// around them, for example UPDATE table SET x = '-Infinity'
//
type Float struct {
	Precision int
}

// Real 4 bytes variable-precision, atleast 6 decimal digits precision
type Real struct {
}

// DoublePrecision 8 bytes variable-precision, atleast 15 decimal digits precision
type DoublePrecision struct {
}

// SmallSerial 2 bytes small autoincrementing integer 1 to 32767
// type names: smallserial and serial2 create smallint
type SmallSerial struct {
}

// Serial 4 bytes autoincrementing integer 1 to 2147483647
// type names: serial and serial4
type Serial struct {
}

// BigSerial 8 bytes large autoincrementing integer 1 to 9223372036854775807
// type names: bigserial and serial8 create bigints
type BigSerial struct {
}

// Postgres Money Types
// https://www.postgresql.org/docs/11/datatype-money.html

// Money 8 bytes currency amount -92233720368547758.08 to +92233720368547758.07
type Money struct {
}

// Character types
// https://www.postgresql.org/docs/11/datatype-character.html
//
// The storage requirement for a short string (up to 126 bytes) is
// 1 byte plus the actual string, which includes the space padding
// in the case of character. Longer strings have 4 bytes of overhead
// instead of 1. Long strings are compressed by the system automatically,
// so the physical requirement on disk might be less. Very long values are
// also stored in background tables so that they do not interfere with rapid
// access to shorter column values. In any case, the longest possible character
// string that can be stored is about 1 GB.

// VarChar type
// character varying(n), varchar(n) variable-length with limit
type VarChar struct {
	Count int // Character Count where n is a positive number
}

// Char type
// character(n), char(n) fixed-length, blank padded
type Char struct {
	Count int // Character Count
}

// Text type
// text variable unlimited lengt
type Text struct {
}

// Byte represents the special "char" type
// char 1 byte single-byte internal type
type Byte struct {
}

// Name is a special internal type for object names
// name 64 bytes internal type for object names
type Name struct {
}

// Binary Types
// https://www.postgresql.org/docs/11/datatype-binary.html

// ByteA 1 or 4 bytes plus the actual binary string variable-length binary string
//
// A binary string is a sequence of octets (or bytes). Binary strings are distinguished from
// character strings in two ways. First, binary strings specifically allow storing octets of
// value zero and other “non-printable” octets (usually, octets outside the decimal range 32 to 126).
// Character strings disallow zero octets, and also disallow any other octet values and sequences of
// octet values that are invalid according to the database's selected character set encoding.
// Second, operations on binary strings process the actual bytes, whereas the processing of
// character strings depends on locale settings. In short, binary strings are appropriate for
// storing data that the programmer thinks of as “raw bytes”, whereas character strings are
// appropriate for storing text.
//
// The bytea type supports two formats for input and output: “hex” format and PostgreSQL's historical
// “escape” format. Both of these are always accepted on input. The output format depends on the
// configuration parameter bytea_output; the default is hex. (Note that the hex format was introduced
// in PostgreSQL 9.0; earlier versions and some tools don't understand it.)
//
type ByteA struct {
}

// Blob type
//
// The SQL standard defines a different binary string type, called BLOB or BINARY LARGE OBJECT.
// The input format is different from bytea, but the provided functions and operators are
// mostly the same.
type Blob struct{}

// Date/Time Types
// https://www.postgresql.org/docs/11/datatype-datetime.html

// Timestamp without a timezone
// timestamp [ (p) ] [ without time zone ] 8 bytes both date and time (no time zone) 4713 BC 294276 AD 1 microsecond
type Timestamp struct {
	Precision int
}

// TimestampZ is a timestamp with a timezone
// timestamp [ (p) ] with time zone, timestampz(p) 8 bytes both date and time, with time zone 4713 BC 294276 AD 1 microsecond
type TimestampZ struct {
	Precision int
}

// Date without a time of day
// date 4 bytes date (no time of day) 4713 BC 5874897 AD 1 day
//
// Acceptable inputs:
//
// 1999-01-08	ISO 8601; January 8 in any mode (recommended format)
// January 8, 1999	unambiguous in any datestyle input mode
// 1/8/1999	January 8 in MDY mode; August 1 in DMY mode
// 1/18/1999	January 18 in MDY mode; rejected in other modes
// 01/02/03	January 2, 2003 in MDY mode; February 1, 2003 in DMY mode; February 3, 2001 in YMD mode
// 1999-Jan-08	January 8 in any mode
// Jan-08-1999	January 8 in any mode
// 08-Jan-1999	January 8 in any mode
// 99-Jan-08	January 8 in YMD mode, else error
// 08-Jan-99	January 8, except error in YMD mode
// Jan-08-99	January 8, except error in YMD mode
// 19990108	ISO 8601; January 8, 1999 in any mode
// 990108	ISO 8601; January 8, 1999 in any mode
// 1999.008	year and day of year
// J2451187	Julian date
// January 8, 99 BC	year 99 BC
type Date struct {
}

// Time without timezone or date
// time [ (p) ] [ without time zone ] 8 bytes time of day (no date) 00:00:00 24:00:00 1 microsecond
//
// Acceptable inputs:
//
// 04:05:06.789	ISO 8601
// 04:05:06	ISO 8601
// 04:05	ISO 8601
// 040506	ISO 8601
// 04:05 AM	same as 04:05; AM does not affect value
// 04:05 PM	same as 16:05; input hour must be <= 12
// 04:05:06.789-8	ISO 8601
// 04:05:06-08:00	ISO 8601
// 04:05-08:00	ISO 8601
// 040506-08	ISO 8601
// 04:05:06 PST	time zone specified by abbreviation
// 2003-04-12 04:05:06 America/New_York	time zone specified by full name
type Time struct {
	Precision int
}

// Acceptable Timezone inputs:
//
// PST	Abbreviation (for Pacific Standard Time)
// America/New_York	Full time zone name
// PST8PDT	POSIX-style time zone specification
// -8:00	ISO-8601 offset for PST
// -800	ISO-8601 offset for PST
// -8	ISO-8601 offset for PST
// zulu	Military abbreviation for UTC
// z	Short form of zulu
//
// select * from pg_timezone_names; to get all the timezone names and their abbreviations
//
// show time zone; to show the databases default timezone

// TimeZ with timezone but no date
// time [ (p) ] with time zone 12 bytes time of day (no date), with time zone 00:00:00+1459 24:00:00-1459 1 microsecond
type TimeZ struct {
	Precision int
}

// Interval is a time interval
// interval [ fields ] [ (p) ]	16 bytes	time interval	-178000000 years	178000000 years	1 microsecond
//
// Interval input syntax looks like this:
//
// [@] quantity unit [quantity unit...] [direction]
// where quantity is a number (possibly signed);
// unit is microsecond, millisecond, second, minute, hour, day, week, month, year, decade, century, millennium, or abbreviations or plurals of these units;
// e.g. '1 day 12 hours 59 min 10 sec 200 years 10 months'
// direction can be ago or empty.
// The at sign (@) is optional noise.
// The amounts of the different units are implicitly added with appropriate sign accounting.
// ago negates all the fields.
//
// Concise input:
// 1-2	SQL standard format: 1 year 2 months
// 3 4:05:06	SQL standard format: 3 days 4 hours 5 minutes 6 seconds
// 1 year 2 months 3 days 4 hours 5 minutes 6 seconds	Traditional Postgres format: 1 year 2 months 3 days 4 hours 5 minutes 6 seconds
// P1Y2M3DT4H5M6S	ISO 8601 “format with designators”: same meaning as above
// P0001-02-03T04:05:06	ISO 8601 “alternative format”: same meaning as above
type Interval struct {
	Precision int
	Fields    IntervalField
}

// IntervalField restricts the set of stored fields.
//
// Possible values:
//
// YEAR
// MONTH
// DAY
// HOUR
// MINUTE
// SECOND
// YEAR TO MONTH
// DAY TO HOUR
// DAY TO MINUTE
// DAY TO SECOND
// HOUR TO MINUTE
// HOUR TO SECOND
// MINUTE TO SECOND
type IntervalField string

// The types abstime and reltime are lower precision types which are used internally.
// You are discouraged from using these types in applications; these internal types
// might disappear in a future release.

// AbsTime type
type AbsTime struct {
}

// Reltime type
type Reltime struct {
}

// Special input types
//
// epoch	date, timestamp	1970-01-01 00:00:00+00 (Unix system time zero)
// infinity	date, timestamp	later than all other time stamps
// -infinity	date, timestamp	earlier than all other time stamps
// now	date, time, timestamp	current transaction's start time
// today	date, timestamp	midnight today
// tomorrow	date, timestamp	midnight tomorrow
// yesterday	date, timestamp	midnight yesterday
// allballs	time	00:00:00.00 UTC

// Postgres output:
// ISO	ISO 8601, SQL standard	1997-12-17 07:37:16-08
//
// ISO 8601 specifies the use of uppercase letter T to separate the date and time.
// PostgreSQL accepts that format on input, but on output it uses a space rather than T,
// as shown above. This is for readability and for consistency with RFC 3339 as well
// as some other database systems.

// Note: show datestyle; to query this information
//
// In the SQL and POSTGRES styles, day appears before month if DMY field ordering has been specified,
// otherwise month appears before day. (See Section 8.5.1 for how this setting also affects
// interpretation of input values.) Table 8.15 shows examples.
//
// datestyle Setting	Input Ordering	Example Output
// SQL, DMY	day/month/year	17/12/1997 15:37:16.00 CET
// SQL, MDY	month/day/year	12/17/1997 07:37:16.00 PST
// Postgres, DMY	day/month/year	Wed 17 Dec 07:37:16 1997 PST

// Boolean Types
// https://www.postgresql.org/docs/11/datatype-boolean.html
// boolean	1 byte	state of true or false
//
// input: TRUE or FALSE
type Boolean struct {
}

// Geometric Types
// https://www.postgresql.org/docs/11/datatype-geometric.html
//
//
// Name	Storage Size	Description	Representation
// point	16 bytes	Point on a plane	(x,y)
// line	32 bytes	Infinite line	{A,B,C}
// lseg	32 bytes	Finite line segment	((x1,y1),(x2,y2))
// box	32 bytes	Rectangular box	((x1,y1),(x2,y2))
// path	16+16n bytes	Closed path (similar to polygon)	((x1,y1),...)
// path	16+16n bytes	Open path	[(x1,y1),...]
// polygon	40+16n bytes	Polygon (similar to closed path)	((x1,y1),...)
// circle	24 bytes	Circle	<(x,y),r> (center point and radius)

// Point type
// point	16 bytes	Point on a plane	(x,y)
//
// where x and y are the respective coordinates, as floating-point numbers.
type Point struct {
	// X float32
	// Y float32
}

// Line type
// line 32 bytes Infinite line {A,B,C}
//
// Lines are represented by the linear equation Ax + By + C = 0, where A and B are not both zero
//
// Also an acceptable input: [ ( x1 , y1 ) , ( x2 , y2 ) ]
type Line struct {
	// A float32
	// B float32
	// C float32
}

// LineSegment type
// lseg	32 bytes	Finite line segment	((x1,y1),(x2,y2))
type LineSegment struct {
}

// Box type
// box	32 bytes	Rectangular box	((x1,y1),(x2,y2))
//
// Boxes are represented by pairs of points that are opposite corners of the box.
// Values of type box are specified using any of the following syntaxes:
type Box struct {
}

// Path type
// path	16+16n bytes	Closed path (similar to polygon)	((x1,y1),...)
// path	16+16n bytes	Open path	[(x1,y1),...]
type Path struct{}

// Polygon type
// polygon	40+16n bytes	Polygon (similar to closed path)	((x1,y1),...)
type Polygon struct{}

// Circle type
// circle	24 bytes	Circle	<(x,y),r> (center point and radius)
type Circle struct{}

//
//  Network Address Types
//
//  https://www.postgresql.org/docs/11/datatype-net-types.html
//  PostgreSQL offers data types to store IPv4, IPv6, and MAC addresses,
//  as shown in Table 8.21. It is better to use these types instead of
//  plain text types to store network addresses, because these types offer
//  input error checking and specialized operators and functions
//
//  Name	Storage Size	Description
//  cidr	7 or 19 bytes	IPv4 and IPv6 networks
//  inet	7 or 19 bytes	IPv4 and IPv6 hosts and networks
//  macaddr	6 bytes	MAC addresses
//  macaddr8	8 bytes	MAC addresses (EUI-64 format)
//
//  cidr Input	cidr Output	abbrev(cidr)
//  192.168.100.128/25	192.168.100.128/25	192.168.100.128/25
//  192.168/24	192.168.0.0/24	192.168.0/24
//  192.168/25	192.168.0.0/25	192.168.0.0/25
//  192.168.1	192.168.1.0/24	192.168.1/24
//  192.168	192.168.0.0/24	192.168.0/24
//  128.1	128.1.0.0/16	128.1/16
//  128	128.0.0.0/16	128.0/16
//  128.1.2	128.1.2.0/24	128.1.2/24
//  10.1.2	10.1.2.0/24	10.1.2/24
//  10.1	10.1.0.0/16	10.1/16
//  10	10.0.0.0/8	10/8
//  10.1.2.3/32	10.1.2.3/32	10.1.2.3/32
//  2001:4f8:3:ba::/64	2001:4f8:3:ba::/64	2001:4f8:3:ba::/64
//  2001:4f8:3:ba:2e0:81ff:fe22:d1f1/128	2001:4f8:3:ba:2e0:81ff:fe22:d1f1/128	2001:4f8:3:ba:2e0:81ff:fe22:d1f1
//  ::ffff:1.2.3.0/120	::ffff:1.2.3.0/120	::ffff:1.2.3/120
//  ::ffff:1.2.3.0/128	::ffff:1.2.3.0/128	::ffff:1.2.3.0/128

// INet type
// inet	7 or 19 bytes	IPv4 and IPv6 hosts and networks
//
// The input format for this type is address/y where address is
// an IPv4 or IPv6 address and y is the number of bits in the netmask.
//
// If the /y portion is missing, the netmask is 32 for IPv4 and 128 for IPv6,
// so the value represents just a single host. On display, the /y portion
// is suppressed if the netmask specifies a single host.
type INet struct{}

// CIDR type
// cidr	7 or 19 bytes	IPv4 and IPv6 networks
//
// The cidr type holds an IPv4 or IPv6 network specification.
// Input and output formats follow Classless Internet Domain Routing
// conventions. The format for specifying networks is address/y where
// address is the network represented as an IPv4 or IPv6 address, and
// y is the number of bits in the netmask. If y is omitted, it is
// calculated using assumptions from the older classful network numbering
// system, except it will be at least large enough to include all of the
// octets written in the input. It is an error to specify a network address
// that has bits set to the right of the specified netmask.
type CIDR struct{}

// MacAddr type
// macaddr	6 bytes	MAC addresses
//
// Input formats:
//
// '08:00:2b:01:02:03'
// '08-00-2b-01-02-03'
// (below not part of the standard)
// '08002b:010203'
// '08002b-010203'
// '0800.2b01.0203'
// '0800-2b01-0203'
// '08002b010203'
type MacAddr struct{}

// Macaddr8 type
//
// The macaddr8 type stores MAC addresses in EUI-64 format,
// known for example from Ethernet card hardware addresses
// (although MAC addresses are used for other purposes as well).
// This type can accept both 6 and 8 byte length MAC addresses
// and stores them in 8 byte length format. MAC addresses given in
// 6 byte format will be stored in 8 byte length format with the 4th
// and 5th bytes set to FF and FE, respectively. Note that IPv6 uses
// a modified EUI-64 format where the 7th bit should be set to one after
// the conversion from EUI-48. The function macaddr8_set7bit is provided
// to make this change. Generally speaking, any input which is comprised
// of pairs of hex digits (on byte boundaries), optionally separated
// consistently by one of ':', '-' or '.', is accepted. The number of hex
// digits must be either 16 (8 bytes) or 12 (6 bytes). Leading and trailing
// whitespace is ignored.
//
// Input formats:
//
// '08:00:2b:01:02:03:04:05'
// '08-00-2b-01-02-03-04-05'
// '08002b:0102030405'
// '08002b-0102030405'
// '0800.2b01.0203.0405'
// '0800-2b01-0203-0405'
// '08002b01:02030405'
// '08002b0102030405'
type Macaddr8 struct{}

// Bit String Types
// https://www.postgresql.org/docs/11/datatype-bit.html
//
// Bit strings are strings of 1's and 0's. They can be used to store or
// visualize bit masks. There are two SQL bit types: bit(n) and bit varying(n),
// where n is a positive integer.
//
// If one explicitly casts a bit-string value to bit(n), it will be truncated
// or zero-padded on the right to be exactly n bits, without raising an error.
// Similarly, if one explicitly casts a bit-string value to bit varying(n),
// it will be truncated on the right if it is more than n bits.

// Bit type data must match the length n exactly; it is an error to attempt
// to store shorter or longer bit strings.
type Bit struct {
	Length int
}

// BitVarying data is of variable length up to the maximum length n;
// longer strings will be rejected. Writing bit without a length is
// equivalent to bit(1), while bit varying without a length
// specification means unlimited length.
type BitVarying struct {
	MaxLength int
}

// Text Search Types
// https://www.postgresql.org/docs/11/datatype-textsearch.html
//
// PostgreSQL provides two data types that are designed to support
// full text search, which is the activity of searching through a
// collection of natural-language documents to locate those that
// best match a query. The tsvector type represents a document in a
// form optimized for text search; the tsquery type similarly represents
// a text query. Chapter 12 provides a detailed explanation of this facility,
// and Section 9.13 summarizes the related functions and operators.

// TSVector type
//
// Value is a sorted list of distinct lexemes, which are words that
// have been normalized to merge different variants of the same word
//
// IMPORTANT: It is important to understand that the tsvector type itself does not
// perform any word normalization; it assumes the words it is given are
// normalized appropriately for the application.
// >>> MATT: this basically means it's a string
//
// Sorting and duplicate-elimination are done automatically during input,
// as shown in this example:
//
// SELECT 'a fat cat sat on a mat and ate a fat rat'::tsvector;
// tsvector
// ----------------------------------------------------
// 'a' 'and' 'ate' 'cat' 'fat' 'mat' 'on' 'rat' 'sat'
//
// Lexemes that have positions can further be labeled with a weight,
// which can be A, B, C, or D. D is the default and hence is not shown on output:
//
// SELECT 'a:1A fat:2B,4C cat:5D'::tsvector;
// ----------------------------
//  'a':1A 'cat':5 'fat':2B,4C
type TSVector struct {
}

// TSQuery type
//
// A tsquery value stores lexemes that are to be searched for, and can
// combine them using the Boolean operators & (AND), | (OR), and ! (NOT),
// as well as the phrase search operator <-> (FOLLOWED BY). There is also a
// variant <N> of the FOLLOWED BY operator, where N is an integer constant
// that specifies the distance between the two lexemes being searched for.
// <-> is equivalent to <1>.
//
// Parentheses can be used to enforce grouping of these operators.
// In the absence of parentheses, ! (NOT) binds most tightly, <-> (FOLLOWED BY)
// next most tightly, then & (AND), with | (OR) binding the least tightly.
//
// SELECT 'fat & rat'::tsquery;
// ---------------
//  'fat' & 'rat'
//
// SELECT 'fat & (rat | cat)'::tsquery;
// ---------------------------
//  'fat' & ( 'rat' | 'cat' )
//
// SELECT 'fat & rat & ! cat'::tsquery;
// ------------------------
//  'fat' & 'rat' & !'cat'
type TSQuery struct {
}

// UUID type
// https://www.postgresql.org/docs/11/datatype-uuid.html
//
// The data type uuid stores Universally Unique Identifiers (UUID) as defined
// by RFC 4122, ISO/IEC 9834-8:2005, and related standards. (Some systems refer
// to this data type as a globally unique identifier, or GUID, instead.)
// This identifier is a 128-bit quantity that is generated by an algorithm
// chosen to make it very unlikely that the same identifier will be generated
// by anyone else in the known universe using the same algorithm. Therefore,
// for distributed systems, these identifiers provide a better uniqueness
// guarantee than sequence generators, which are only unique within a single
// database.
//
// A UUID is written as a sequence of lower-case hexadecimal digits, in several
// groups separated by hyphens, specifically a group of 8 digits followed by
// three groups of 4 digits followed by a group of 12 digits, for a total of
// 32 digits representing the 128 bits. An example of a UUID in this standard
// form is: a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11
//
// PostgreSQL provides storage and comparison functions for UUIDs, but the
// core database does not include any function for generating UUIDs,
// because no single algorithm is well suited for every application.
// The uuid-ossp module provides functions that implement several standard
// algorithms. The pgcrypto module also provides a generation function for
// random UUIDs. Alternatively, UUIDs could be generated by client applications
// or other libraries invoked through a server-side function.
type UUID struct{}

// XML Type
// https://www.postgresql.org/docs/11/datatype-xml.html
//
// The xml data type can be used to store XML data. Its advantage over
// storing XML data in a text field is that it checks the input values
// for well-formedness, and there are support functions to perform
// type-safe operations on it; see Section 9.14. Use of this data type
// requires the installation to have been built with configure --with-libxml.
//
// The xml type can store well-formed “documents”, as defined by the XML
// standard, as well as “content” fragments, which are defined by the
// production XMLDecl? content in the XML standard. Roughly, this means
// that content fragments can have more than one top-level element or
// character node. The expression xmlvalue IS DOCUMENT can be used to
// evaluate whether a particular xml value is a full document or only a
// content fragment.
//
// To produce a value of type xml from character data, use the function xmlparse:
//
// XMLPARSE (DOCUMENT '<?xml version="1.0"?><book><title>Manual</title><chapter>...</chapter></book>')
// XMLPARSE (CONTENT 'abc<foo>bar</foo><bar>foo</bar>')
type XML struct{}

// JSON type
// https://www.postgresql.org/docs/11/datatype-json.html
//
// JSON data types are for storing JSON (JavaScript Object Notation) data,
// as specified in RFC 7159. Such data can also be stored as text, but the
// JSON data types have the advantage of enforcing that each stored value
// is valid according to the JSON rules. There are also assorted JSON-specific
// functions and operators available for data stored in these data types;
//
// There are two JSON data types: json and jsonb. They accept almost identical
// sets of values as input. The major practical difference is one of efficiency.
// The json data type stores an exact copy of the input text, which processing
// functions must reparse on each execution; while jsonb data is stored in a
// decomposed binary format that makes it slightly slower to input due to added
// conversion overhead, but significantly faster to process, since no reparsing
// is needed. jsonb also supports indexing, which can be a significant advantage.
//
// Because the json type stores an exact copy of the input text, it will preserve
// semantically-insignificant white space between tokens, as well as the order of
// keys within JSON objects. Also, if a JSON object within the value contains the
// same key more than once, all the key/value pairs are kept. (The processing
// functions consider the last value as the operative one.) By contrast, jsonb
// does not preserve white space, does not preserve the order of object keys,
// and does not keep duplicate object keys. If duplicate keys are specified in
// the input, only the last value is kept.
//
// In general, most applications should prefer to store JSON data as jsonb,
// unless there are quite specialized needs, such as legacy assumptions
// about ordering of object keys.
//
// JSON primitive type	PostgreSQL type	Notes
// string	text	\u0000 is disallowed, as are non-ASCII Unicode escapes if database encoding is not UTF8
// number	numeric	NaN and infinity values are disallowed
// boolean	boolean	Only lowercase true and false spellings are accepted
// null	(none)	SQL NULL is a different concept

// JSON type
type JSON struct{}

// JSONB type is the binary JSON type
type JSONB struct{}

// Array Types
// https://www.postgresql.org/docs/11/arrays.html
//
// PostgreSQL allows columns of a table to be defined as variable-length
// multidimensional arrays. Arrays of any built-in or user-defined base
// type, enum type, composite type, range type, or domain can be created.
//
// Example:
//
// CREATE TABLE sal_emp (
//   name            text,
//   pay_by_quarter  integer[],
//   schedule        text[][]
// );
//
// An array data type is named by appending square brackets ([]) to the data
// type name of the array elements. The above command will create a table
// named sal_emp with a column of type text (name), a one-dimensional array
// of type integer (pay_by_quarter), which represents the employee's salary
// by quarter, and a two-dimensional array of text (schedule), which represents
// the employee's weekly schedule.
//
// The syntax for CREATE TABLE allows the exact size of arrays to be specified, for example:
//
// CREATE TABLE tictactoe (
//   squares   integer[3][3]
// );
//
// An alternative syntax, which conforms to the SQL standard by using the keyword
// ARRAY, can be used for one-dimensional arrays. pay_by_quarter could have been
// defined as:
//
// pay_by_quarter  integer ARRAY[4], // 1 dimension array with 4 ints
// pay_by_quarter  integer ARRAY,    // no array size
//
// Input: '{ val1 delim val2 delim ... }'
//
// where delim is the delimiter character for the type, as recorded in its
// pg_type entry. Among the standard data types provided in the PostgreSQL
// distribution, all use a comma (,), except for type box which uses a
// semicolon (;). Each val is either a constant of the array element type, or
// a subarray. An example of an array constant is:
//
//    '{{1,2,3},{4,5,6},{7,8,9}}'
//
// This constant is a two-dimensional, 3-by-3 array consisting of three
// subarrays of integers.
//
// Arrays are not sets; searching for specific array elements can be a sign
// of database misdesign. Consider using a separate table with a row for each
// item that would be an array element. This will be easier to search, and
// is likely to scale better for a large number of elements.
//
// The ARRAY constructor syntax (see Section 4.2.12) is often easier to work
// with than the array-literal syntax when writing array values in SQL commands.
// In ARRAY, individual element values are written the same way they would be
// written when not members of an array.
type Array struct {
	Type DataType
}

// Composite Types
// https://www.postgresql.org/docs/11/rowtypes.html
//
// A composite type represents the structure of a row or record; it is essentially
// just a list of field names and their data types. PostgreSQL allows composite
// types to be used in many of the same ways that simple types can be used.
// For example, a column of a table can be declared to be of a composite type.
//
// CREATE TYPE complex AS (
//   r       double precision,
//   i       double precision
// );
//
// CREATE TYPE inventory_item AS (
//   name            text,
//   supplier_id     integer,
//   price           numeric
// );
//
// Usage:
//
// CREATE TABLE on_hand (
//   item      inventory_item,
//   count     integer
// );
//
// INSERT INTO on_hand VALUES (ROW('fuzzy dice', 42, 1.99), 1000);
//
// Inputs: '( val1 , val2 , ... )'
//
// The ROW expression syntax can also be used to construct composite values.
// In most cases this is considerably simpler to use than the string-literal
// syntax since you don't have to worry about multiple layers of quoting.
// We already used this method above:
//
// ROW('fuzzy dice', 42, 1.99)
// ROW('', 42, NULL)
type Composite struct {
	Fields []CompositeField
}

// CompositeField type
type CompositeField struct {
	Name string
	Type DataType
}

// Range Types
// https://www.postgresql.org/docs/11/rangetypes.html
//
// Range types are data types representing a range of values of some
// element type (called the range's subtype). For instance, ranges
// of timestamp might be used to represent the ranges of time that a
// meeting room is reserved. In this case the data type is tsrange
// (short for “timestamp range”), and timestamp is the subtype. The subtype
// must have a total order so that it is well-defined whether element values
// are within, before, or after a range of values.
//
// Range types are useful because they represent many element values in
// a single range value, and because concepts such as overlapping ranges
// can be expressed clearly. The use of time and date ranges for scheduling
// purposes is the clearest example; but price ranges, measurement ranges
// from an instrument, and so forth can also be useful.
//
// You can define your own range types; see CREATE TYPE for more information.
//
//   CREATE TYPE floatrange AS RANGE (
//     subtype = float8,
//     subtype_diff = float8mi
//   );
//   SELECT '[1.234, 5.678]'::floatrange;
//
// The input for a range value must follow one of the following patterns:
//
//   (lower-bound,upper-bound)
//   (lower-bound,upper-bound]
//   [lower-bound,upper-bound)
//   [lower-bound,upper-bound]
//   empty
//
//   CREATE TABLE reservation (room int, during tsrange);
//   INSERT INTO reservation VALUES
//       (1108, '[2010-01-01 14:30, 2010-01-01 15:30)');

// Int4Range – Range of integer
type Int4Range struct {
}

// Int8Range – Range of bigint
type Int8Range struct {
}

// NumRange — Range of numeric
type NumRange struct {
}

// TSRange — Range of timestamp without time zone
type TSRange struct {
}

// TSTZRange — Range of timestamp with time zone
type TSTZRange struct {
}

// DateRange — Range of date
type DateRange struct {
}

// Domain Types
// https://www.postgresql.org/docs/11/domains.html
//
// A domain is a user-defined data type that is based on another underlying type.
// Optionally, it can have constraints that restrict its valid values to a subset
// of what the underlying type would allow. Otherwise it behaves like the
// underlying type — for example, any operator or function that can be applied
// to the underlying type will work on the domain type. The underlying type can
// be any built-in or user-defined base type, enum type, array type, composite
// type, range type, or another domain.
//
// CREATE DOMAIN posint AS integer CHECK (VALUE > 0);
// CREATE TABLE mytable (id posint);
// INSERT INTO mytable VALUES(1);   -- works
// INSERT INTO mytable VALUES(-1);  -- fails
type Domain struct {
	Type DataType
}

// Object Identifier Types
// https://www.postgresql.org/docs/11/datatype-oid.html
//
// Object identifiers (OIDs) are used internally by PostgreSQL as primary
// keys for various system tables. OIDs are not added to user-created tables,
// unless WITH OIDS is specified when the table is created, or the
// default_with_oids configuration variable is enabled. Type oid represents
// an object identifier.
//
// There are also several alias types for oid: regproc,
// regprocedure, regoper, regoperator, regclass, regtype, regrole, regnamespace,
// regconfig, and regdictionary. The OID alias types have no operations of their
// own except for specialized input and output routines.
//
// The oid type is currently implemented as an unsigned four-byte integer.
// Therefore, it is not large enough to provide database-wide uniqueness
// in large databases, or even in large individual tables. So, using a
// user-created table's OID column as a primary key is discouraged.
// OIDs are best used only for references to system tables.
//
// OID types:
//
// Name	References	Description	Value Example
// oid	any	numeric object identifier	564182
// regproc	pg_proc	function name	sum
// regprocedure	pg_proc	function with argument types	sum(int4)
// regoper	pg_operator	operator name	+
// regoperator	pg_operator	operator with argument types	*(integer,integer) or -(NONE,integer)
// regclass	pg_class	relation name	pg_type
// regtype	pg_type	data type name	integer
// regrole	pg_authid	role name	smithee
// regnamespace	pg_namespace	namespace name	pg_catalog
// regconfig	pg_ts_config	text search configuration	english
// regdictionary	pg_ts_dict	text search dictionary	simple
//
// Another identifier type used by the system is xid, or transaction
// (abbreviated xact) identifier. This is the data type of the system
// columns xmin and xmax. Transaction identifiers are 32-bit quantities.
//
// A third identifier type used by the system is cid, or command identifier.
// This is the data type of the system columns cmin and cmax. Command
// identifiers are also 32-bit quantities.
//
// A final identifier type used by the system is tid, or tuple identifier
// (row identifier). This is the data type of the system column ctid.
// A tuple ID is a pair (block number, tuple index within block) that
// identifies the physical location of the row within its table.

// pg_lsn type
// https://www.postgresql.org/docs/11/datatype-pg-lsn.html
//
// The pg_lsn data type can be used to store LSN (Log Sequence Number)
// data which is a pointer to a location in the WAL. This type is a
// representation of XLogRecPtr and an internal system type of PostgreSQL.
//
// Internally, an LSN is a 64-bit integer, representing a byte position in
// the write-ahead log stream. It is printed as two hexadecimal numbers of
// up to 8 digits each, separated by a slash; for example, 16/B374D848.
// The pg_lsn type supports the standard comparison operators, like = and >.
// Two LSNs can be subtracted using the - operator; the result is the number
// of bytes separating those write-ahead log locations.

// Pseudo-Types
// https://www.postgresql.org/docs/11/datatype-pseudo.html
//
// The PostgreSQL type system contains a number of special-purpose entries
// that are collectively called pseudo-types.
//
// A pseudo-type cannot be used as a column data type, but it can be
// used to declare a function's argument or result type. Each of the
// available pseudo-types is useful in situations where a function's
// behavior does not correspond to simply taking or returning a value
// of a specific SQL data type.
//
// Name	Description
// any	Indicates that a function accepts any input data type.
// anyelement	Indicates that a function accepts any data type (see Section 38.2.5).
// anyarray	Indicates that a function accepts any array data type (see Section 38.2.5).
// anynonarray	Indicates that a function accepts any non-array data type (see Section 38.2.5).
// anyenum	Indicates that a function accepts any enum data type (see Section 38.2.5 and Section 8.7).
// anyrange	Indicates that a function accepts any range data type (see Section 38.2.5 and Section 8.17).
// cstring	Indicates that a function accepts or returns a null-terminated C string.
// internal	Indicates that a function accepts or returns a server-internal data type.
// language_handler	A procedural language call handler is declared to return language_handler.
// fdw_handler	A foreign-data wrapper handler is declared to return fdw_handler.
// index_am_handler	An index access method handler is declared to return index_am_handler.
// tsm_handler	A tablesample method handler is declared to return tsm_handler.
// record	Identifies a function taking or returning an unspecified row type.
// trigger	A trigger function is declared to return trigger.
// event_trigger	An event trigger function is declared to return event_trigger.
// pg_ddl_command	Identifies a representation of DDL commands that is available to event triggers.
// void	Indicates that a function returns no value.
// unknown	Identifies a not-yet-resolved type, e.g. of an undecorated string literal.
// opaque	An obsolete type name that formerly served many of the above purposes.
