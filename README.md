libgosysstat
============

CPU, Memory and Disk stat tool written in Go


Uses
====

A library for reading:

 * /proc/stat
 * /proc/meminfo

at a set interval.


Mental Model
============

    +-----------+
    | processor |- StatProcessor(statProcessor, chan)
    +-----------+
        |-- Process() -> []core.Stat{}
    +-----------------+
    | ResettingReader | Open(reader.DataSource)
    +-----------------+ 
        |-- Read() -> [][]string
    +---------+
    | io.File |
    +---------+
        |-- Seek()
        \-- Read()


TODO
====

 + disk stats (% used, io rates)
 + net stats (kb/s?)
 + /proc/<pid>/stat
