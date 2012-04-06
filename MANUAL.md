    `import "github.com/rug-compling/alpinocorpus-go/alpinocorpus"`

    [Overview][1]

    [Index][2]

## Overview ▹

## Overview ▾

Package alpinocorpus provides a reader and a writer for Alpino corpora

## Index

    [Constants][3]

    [func WriterAvailable(writertype writerType) bool][4]

    [type Entries][5]

        [func (it *Entries) Close()][6]

        [func (it *Entries) Contents() string][7]

        [func (it *Entries) Name() string][8]

        [func (it *Entries) Next() bool][9]

    [type Reader][10]

        [func NewReader(filename string) (*Reader, error)][11]

        [func (r *Reader) Close()][12]

        [func (r *Reader) Get(entry string) (string, error)][13]

        [func (r *Reader) GetAll() (*Entries, error)][14]

        [func (r *Reader) GetMod(entry, markerQuery, markerAttr, markerValue
string) (string, error)][15]

        [func (r *Reader) Len() int][16]

        [func (r *Reader) Query(query string) (*Entries, error)][17]

        [func (r *Reader) QueryMod(query, markerQuery, markerAttr,
markerValue, stylesheet string) (*Entries, error)][18]

        [func (r *Reader) ValidQuery(query string) bool][19]

    [type Writer][20]

        [func NewWriter(filename string, overwrite bool) (*Writer, error)][21]

        [func NewWriterType(filename string, overwrite bool, writertype
writerType) (*Writer, error)][22]

        [func (w *Writer) Close()][23]

        [func (w *Writer) Write(name, contents string) error][24]

        [func (w *Writer) WriteCorpus(r *Reader, failsafe bool) error][25]

#### Package files

[Reader.go][26] [Writer.go][27]

## Constants


    const (

        Compact writerType = iota

        Dbxml

    )

## func [WriterAvailable][28]


    func WriterAvailable(writertype writerType) bool

Check whether a particular writer type is available. Currently, the only valid
types are Compact and Dbxml.

## type [Entries][29]


    type Entries struct {

        // contains filtered or unexported fields

    }

### func (*Entries) [Close][30]


    func (it *Entries) Close()

### func (*Entries) [Contents][31]


    func (it *Entries) Contents() string

### func (*Entries) [Name][32]


    func (it *Entries) Name() string

### func (*Entries) [Next][33]


    func (it *Entries) Next() bool

## type [Reader][34]


    type Reader struct {

        // contains filtered or unexported fields

    }

### func [NewReader][35]


    func NewReader(filename string) (*Reader, error)

NewReader() opens an Alpino corpus for reading

### func (*Reader) [Close][36]


    func (r *Reader) Close()

Close() closes the corpus

### func (*Reader) [Get][37]


    func (r *Reader) Get(entry string) (string, error)

Get() returns the contents of an entry given its label

### func (*Reader) [GetAll][38]


    func (r *Reader) GetAll() (*Entries, error)

GetAll() gives access to all entries in the corpus

Example usage:


    entries, error := reader.GetAll()

    if error != nil {

        log.Fatal(error)

    }

    for entries.Next() {

        fmt.Println(entries.Name(), ": ", entries.Contents())

    }

    entries.Close()


Don't forget to call Close() at the end!

### func (*Reader) [GetMod][39]


    func (r *Reader) GetMod(entry, markerQuery, markerAttr, markerValue
string) (string, error)

GetMod() returns the contents of an entry given its label, with some items
marked

### func (*Reader) [Len][40]


    func (r *Reader) Len() int

Len() returns the number of entries in the corpus

### func (*Reader) [Query][41]


    func (r *Reader) Query(query string) (*Entries, error)

Query() gives access to the names of all entries that match a certain query

See GetAll() for an example of how to use the result

### func (*Reader) [QueryMod][42]


    func (r *Reader) QueryMod(query, markerQuery, markerAttr, markerValue,
stylesheet string) (*Entries, error)

QueryMod() gives access to all entries that match a certain query, using a
stylesheet to modify the output

See GetAll() for an example of how to use the result

### func (*Reader) [ValidQuery][43]


    func (r *Reader) ValidQuery(query string) bool

ValidQuery() checks if a query is well-formed

## type [Writer][44]


    type Writer struct {

        // contains filtered or unexported fields

    }

### func [NewWriter][45]


    func NewWriter(filename string, overwrite bool) (*Writer, error)

NewWriter() opens an Alpino corpus for writing. The corpus is of type Dbxml.

### func [NewWriterType][46]


    func NewWriterType(filename string, overwrite bool, writertype writerType)
(*Writer, error)

NewWriterType() opens an Alpino corpus for writing. The type of corpus is
specified in the third argument. Currently, the only valid types are Compact
and Dbxml.

### func (*Writer) [Close][47]


    func (w *Writer) Close()

Close() closes the corpus

### func (*Writer) [Write][48]


    func (w *Writer) Write(name, contents string) error

### func (*Writer) [WriteCorpus][49]


    func (w *Writer) WriteCorpus(r *Reader, failsafe bool) error

   [1]: #overview

   [2]: #index

   [3]: #constants

   [4]: #WriterAvailable

   [5]: #Entries

   [6]: #Entries.Close

   [7]: #Entries.Contents

   [8]: #Entries.Name

   [9]: #Entries.Next

   [10]: #Reader

   [11]: #NewReader

   [12]: #Reader.Close

   [13]: #Reader.Get

   [14]: #Reader.GetAll

   [15]: #Reader.GetMod

   [16]: #Reader.Len

   [17]: #Reader.Query

   [18]: #Reader.QueryMod

   [19]: #Reader.ValidQuery

   [20]: #Writer

   [21]: #NewWriter

   [22]: #NewWriterType

   [23]: #Writer.Close

   [24]: #Writer.Write

   [25]: #Writer.WriteCorpus

   [26]: /target/Reader.go

   [27]: /target/Writer.go

   [28]: /target/Writer.go?s=2584:2632#L112

   [29]: /target/Reader.go?s=346:510#L14

   [30]: /target/Reader.go?s=1093:1119#L52

   [31]: /target/Reader.go?s=912:948#L42

   [32]: /target/Reader.go?s=858:890#L38

   [33]: /target/Reader.go?s=512:542#L23

   [34]: /target/Reader.go?s=1201:1298#L59

   [35]: /target/Reader.go?s=1350:1398#L66

   [36]: /target/Reader.go?s=2110:2134#L100

   [37]: /target/Reader.go?s=2263:2313#L108

   [38]: /target/Reader.go?s=2914:2957#L136

   [39]: /target/Reader.go?s=5466:5557#L232

   [40]: /target/Reader.go?s=1981:2007#L92

   [41]: /target/Reader.go?s=3617:3671#L164

   [42]: /target/Reader.go?s=4147:4254#L183

   [43]: /target/Reader.go?s=3249:3295#L149

   [44]: /target/Writer.go?s=153:250#L5

   [45]: /target/Writer.go?s=512:576#L26

   [46]: /target/Writer.go?s=1133:1224#L47

   [47]: /target/Writer.go?s=2244:2268#L96

   [48]: /target/Writer.go?s=1615:1666#L65

   [49]: /target/Writer.go?s=1947:2007#L80

