PACKAGE package alpinocorpus import "github.com/rug-compling/alpinocorpus-
go/alpinocorpus" Package alpinocorpus provides a reader and a writer for
Alpino corpora CONSTANTS const ( Compact writerType = iota Dbxml ) FUNCTIONS
func WriterAvailable(writertype writerType) bool Check whether a particular
writer type is available. Currently, the only valid types are Compact and
Dbxml. TYPES type Entries struct { // contains filtered or unexported fields }
func (it *Entries) Close() func (it *Entries) Contents() string func (it
*Entries) Name() string func (it *Entries) Next() bool type Reader struct { //
contains filtered or unexported fields } func NewReader(filename string)
(*Reader, error) NewReader() opens an Alpino corpus for reading func (r
*Reader) Close() Close() closes the corpus func (r *Reader) Get(entry string)
(string, error) Get() returns the contents of an entry given its label func (r
*Reader) GetAll() (*Entries, error) GetAll() gives access to all entries in
the corpus Example usage: entries, error := reader.GetAll() if error != nil {
log.Fatal(error) } for entries.Next() { fmt.Println(entries.Name(), ": ",
entries.Contents()) } entries.Close() Don't forget to call Close() at the end!
func (r *Reader) GetMod(entry, markerQuery, markerAttr, markerValue string)
(string, error) GetMod() returns the contents of an entry given its label,
with some items marked func (r *Reader) Len() int Len() returns the number of
entries in the corpus func (r *Reader) Query(query string) (*Entries, error)
Query() gives access to the names of all entries that match a certain query
See GetAll() for an example of how to use the result func (r *Reader)
QueryMod(query, markerQuery, markerAttr, markerValue, stylesheet string)
(*Entries, error) QueryMod() gives access to all entries that match a certain
query, using a stylesheet to modify the output See GetAll() for an example of
how to use the result func (r *Reader) ValidQuery(query string) bool
ValidQuery() checks if a query is well-formed type Writer struct { // contains
filtered or unexported fields } func NewWriter(filename string, overwrite
bool) (*Writer, error) NewWriter() opens an Alpino corpus for writing. The
corpus is of type Dbxml. func NewWriterType(filename string, overwrite bool,
writertype writerType) (*Writer, error) NewWriterType() opens an Alpino corpus
for writing. The type of corpus is specified in the third argument. Currently,
the only valid types are Compact and Dbxml. func (w *Writer) Close() Close()
closes the corpus func (w *Writer) Write(name, contents string) error func (w
*Writer) WriteCorpus(r *Reader, failsafe bool) error

