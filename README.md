gosysstat
=========

CPU, Memory and Disk stat tool written in Go

Uses
====

As a library that you can embed

A CLI which reports stats to your console

Mental Model
============

     +------+
     | main |
     +------+
        |
        |-- cpuFh = os.Open("/proc/stat")
        |-- core.StatProcessor(cpu.NewProcessor(cpuFh), chan)
        |
        |-- memFh = os.Open("/proc/meminfo")
        |-- core.StatProcessor(ram.NewProcessor(memFh), chan)
        |
        |-- core.StatProcessor(dsk.NewProcessor(), chan)
        |            // disk io?, disk storage?
        |
        |-- core.StatProcessor(net.NewProcessor(), chan)
        |            // does this use the /proc fs?
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

