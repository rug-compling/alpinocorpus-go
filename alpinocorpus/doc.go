/*

Package alpinocorpus provides a reader and a writer for Alpino corpora [corpus linguistics].

Example usage:

Opening a corpus for reading:

    reader, error := alpinocorpus.NewReader(filename)
    defer reader.Close()  // to free resources
    if error != nil {
        log.Fatalln(error)
    }

Getting all entries from the corpus:

    entries, error := reader.GetAll()
    if error != nil {
        log.Fatalln(error)
    }

Or, getting al entries that match some query:

    entries, error := reader.Query("//node[@root=\"fiets\"]")
    if error != nil {
        log.Fatalln(error)
    }

And then, this:

    for key := range entries.Keys() {
        fmt.Println(key)
    }

Or this:

    for value := range entries.Values() {
        fmt.Println(value)
    }

Or this:

    for pair := range entries.KeysValues() {
        fmt.Println(pair.Key, ": ", pair.Value)
    }

After one of these, the entries are no longer accessible.

IMPORTANT:

To end an iteration prematurely, use Break(), so all resources are cleaned up

Wrong:

    for key := range entries.Keys() {
        fmt.Println(key)
        if key == somekey {
            break
        }
    }

Right:

    for key := range entries.Keys() {
        fmt.Println(key)
        if key == somekey {
            entries.Break()
        }
    }

To break out of the middle, use Break() followed by break:

    for key := range entries.Keys() {
        fmt.Println(key)
        if key == somekey {
            entries.Break()
            break
        }
        moreStuff()
    }

*/
package alpinocorpus
