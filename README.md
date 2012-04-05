## Dependences

 * [alpinocorpus](https://github.com/rug-compling/alpinocorpus)
 * [libexslt](http://xmlsoft.org/xslt/EXSLT/)

## Install

  go get -u github.com/rug-compling/alpinocorpus-go/alpinocorpus

## Usage

<!--
	Copyright 2009 The Go Authors. All rights reserved.
	Use of this source code is governed by a BSD-style
	license that can be found in the LICENSE file.
-->

	
		<div id="short-nav">
			<dl>
			<dd><code>import "github.com/rug-compling/alpinocorpus-go/alpinocorpus"</code></dd>
			</dl>
			<dl>
			<dd><a href="#overview" class="overviewLink">Overview</a></dd>
			<dd><a href="#index">Index</a></dd>
			
			
			</dl>
		</div>
		<!-- The package's Name is printed as title by the top-level template -->
		<div id="overview" class="toggleVisible">
			<div class="collapsed">
				<h2 class="toggleButton" title="Click to show Overview section">Overview ▹</h2>
			</div>
			<div class="expanded">
				<h2 class="toggleButton" title="Click to hide Overview section">Overview ▾</h2>
				<p>
Package alpinocorpus provides a reader and a writer for Alpino corpora
</p>

			</div>
		</div>
		
	
		<h2 id="index">Index</h2>
		<!-- Table of contents for API; must be named manual-nav to turn off auto nav. -->
		<div id="manual-nav">
			<dl>
			
				<dd><a href="#constants">Constants</a></dd>
			
			
			
				
				<dd><a href="#WriterAvailable">func WriterAvailable(writertype writerType) bool</a></dd>
			
			
				
				<dd><a href="#Entries">type Entries</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Entries.Close">func (it *Entries) Close()</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Entries.Contents">func (it *Entries) Contents() string</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Entries.Name">func (it *Entries) Name() string</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Entries.Next">func (it *Entries) Next() bool</a></dd>
				
			
				
				<dd><a href="#Reader">type Reader</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NewReader">func NewReader(filename string) (*Reader, error)</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.Close">func (r *Reader) Close()</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.Get">func (r *Reader) Get(entry string) (string, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.GetAll">func (r *Reader) GetAll() (*Entries, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.GetMod">func (r *Reader) GetMod(entry, markerQuery, markerAttr, markerValue string) (string, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.Len">func (r *Reader) Len() int</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.Query">func (r *Reader) Query(query string) (*Entries, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.QueryMod">func (r *Reader) QueryMod(query, markerQuery, markerAttr, markerValue, stylesheet string) (*Entries, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Reader.ValidQuery">func (r *Reader) ValidQuery(query string) bool</a></dd>
				
			
				
				<dd><a href="#Writer">type Writer</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NewWriter">func NewWriter(filename string, overwrite bool) (*Writer, error)</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#NewWriterType">func NewWriterType(filename string, overwrite bool, writertype writerType) (*Writer, error)</a></dd>
				
				
					
					<dd>&nbsp; &nbsp; <a href="#Writer.Close">func (w *Writer) Close()</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Writer.Write">func (w *Writer) Write(name, contents string) error</a></dd>
				
					
					<dd>&nbsp; &nbsp; <a href="#Writer.WriteCorpus">func (w *Writer) WriteCorpus(r *Reader, failsafe bool) error</a></dd>
				
			
			
		</dl>

		

		
			<h4>Package files</h4>
			<p>
			<span style="font-size:90%">
			
				<a href="/target/Reader.go">Reader.go</a>
			
				<a href="/target/Writer.go">Writer.go</a>
			
			</span>
			</p>
		
	
		
			<h2 id="constants">Constants</h2>
			
				<pre>const (
    Compact writerType = iota
    Dbxml
)</pre>
				
			
		
		
		
			
			
			<h2 id="WriterAvailable">func <a href="/target/Writer.go?s=2584:2632#L112">WriterAvailable</a></h2>
			<pre>func WriterAvailable(writertype writerType) bool</pre>
			<p>
Check whether a particular writer type is available.
Currently, the only valid types are Compact and Dbxml.
</p>

			
		
		
			
			
			<h2 id="Entries">type <a href="/target/Reader.go?s=346:510#L14">Entries</a></h2>
			<pre>type Entries struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre>
			

			

			

			

			

			
				
				<h3 id="Entries.Close">func (*Entries) <a href="/target/Reader.go?s=1093:1119#L52">Close</a></h3>
				<pre>func (it *Entries) Close()</pre>
				
				
				
			
				
				<h3 id="Entries.Contents">func (*Entries) <a href="/target/Reader.go?s=912:948#L42">Contents</a></h3>
				<pre>func (it *Entries) Contents() string</pre>
				
				
				
			
				
				<h3 id="Entries.Name">func (*Entries) <a href="/target/Reader.go?s=858:890#L38">Name</a></h3>
				<pre>func (it *Entries) Name() string</pre>
				
				
				
			
				
				<h3 id="Entries.Next">func (*Entries) <a href="/target/Reader.go?s=512:542#L23">Next</a></h3>
				<pre>func (it *Entries) Next() bool</pre>
				
				
				
			
		
			
			
			<h2 id="Reader">type <a href="/target/Reader.go?s=1201:1298#L59">Reader</a></h2>
			<pre>type Reader struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre>
			

			

			

			

			
				
				<h3 id="NewReader">func <a href="/target/Reader.go?s=1350:1398#L66">NewReader</a></h3>
				<pre>func NewReader(filename string) (*Reader, error)</pre>
				<p>
NewReader() opens an Alpino corpus for reading
</p>

				
			

			
				
				<h3 id="Reader.Close">func (*Reader) <a href="/target/Reader.go?s=2110:2134#L100">Close</a></h3>
				<pre>func (r *Reader) Close()</pre>
				<p>
Close() closes the corpus
</p>

				
				
			
				
				<h3 id="Reader.Get">func (*Reader) <a href="/target/Reader.go?s=2263:2313#L108">Get</a></h3>
				<pre>func (r *Reader) Get(entry string) (string, error)</pre>
				<p>
Get() returns the contents of an entry given its label
</p>

				
				
			
				
				<h3 id="Reader.GetAll">func (*Reader) <a href="/target/Reader.go?s=2914:2957#L136">GetAll</a></h3>
				<pre>func (r *Reader) GetAll() (*Entries, error)</pre>
				<p>
GetAll() gives access to all entries in the corpus
</p>
<p>
Example usage:
</p>
<pre>entries, error := reader.GetAll()
if error != nil {
    log.Fatal(error)
}
for entries.Next() {
    fmt.Println(entries.Name(), &#34;: &#34;, entries.Contents())
}
entries.Close()
</pre>
<p>
Don&#39;t forget to call Close() at the end!
</p>

				
				
			
				
				<h3 id="Reader.GetMod">func (*Reader) <a href="/target/Reader.go?s=5466:5557#L232">GetMod</a></h3>
				<pre>func (r *Reader) GetMod(entry, markerQuery, markerAttr, markerValue string) (string, error)</pre>
				<p>
GetMod() returns the contents of an entry given its label, with some items marked
</p>

				
				
			
				
				<h3 id="Reader.Len">func (*Reader) <a href="/target/Reader.go?s=1981:2007#L92">Len</a></h3>
				<pre>func (r *Reader) Len() int</pre>
				<p>
Len() returns the number of entries in the corpus
</p>

				
				
			
				
				<h3 id="Reader.Query">func (*Reader) <a href="/target/Reader.go?s=3617:3671#L164">Query</a></h3>
				<pre>func (r *Reader) Query(query string) (*Entries, error)</pre>
				<p>
Query() gives access to the names of all entries that match a certain query
</p>
<p>
See GetAll() for an example of how to use the result
</p>

				
				
			
				
				<h3 id="Reader.QueryMod">func (*Reader) <a href="/target/Reader.go?s=4147:4254#L183">QueryMod</a></h3>
				<pre>func (r *Reader) QueryMod(query, markerQuery, markerAttr, markerValue, stylesheet string) (*Entries, error)</pre>
				<p>
QueryMod() gives access to all entries that match a certain query, using a stylesheet to modify the output
</p>
<p>
See GetAll() for an example of how to use the result
</p>

				
				
			
				
				<h3 id="Reader.ValidQuery">func (*Reader) <a href="/target/Reader.go?s=3249:3295#L149">ValidQuery</a></h3>
				<pre>func (r *Reader) ValidQuery(query string) bool</pre>
				<p>
ValidQuery() checks if a query is well-formed
</p>

				
				
			
		
			
			
			<h2 id="Writer">type <a href="/target/Writer.go?s=153:250#L5">Writer</a></h2>
			<pre>type Writer struct {
    <span class="comment">// contains filtered or unexported fields</span>
}</pre>
			

			

			

			

			
				
				<h3 id="NewWriter">func <a href="/target/Writer.go?s=512:576#L26">NewWriter</a></h3>
				<pre>func NewWriter(filename string, overwrite bool) (*Writer, error)</pre>
				<p>
NewWriter() opens an Alpino corpus for writing.
The corpus is of type Dbxml.
</p>

				
			
				
				<h3 id="NewWriterType">func <a href="/target/Writer.go?s=1133:1224#L47">NewWriterType</a></h3>
				<pre>func NewWriterType(filename string, overwrite bool, writertype writerType) (*Writer, error)</pre>
				<p>
NewWriterType() opens an Alpino corpus for writing.
The type of corpus is specified in the third argument.
Currently, the only valid types are Compact and Dbxml.
</p>

				
			

			
				
				<h3 id="Writer.Close">func (*Writer) <a href="/target/Writer.go?s=2244:2268#L96">Close</a></h3>
				<pre>func (w *Writer) Close()</pre>
				<p>
Close() closes the corpus
</p>

				
				
			
				
				<h3 id="Writer.Write">func (*Writer) <a href="/target/Writer.go?s=1615:1666#L65">Write</a></h3>
				<pre>func (w *Writer) Write(name, contents string) error</pre>
				
				
				
			
				
				<h3 id="Writer.WriteCorpus">func (*Writer) <a href="/target/Writer.go?s=1947:2007#L80">WriteCorpus</a></h3>
				<pre>func (w *Writer) WriteCorpus(r *Reader, failsafe bool) error</pre>
				
				
				
			
		
		</div>
	

	







