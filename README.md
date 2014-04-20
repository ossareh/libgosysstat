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
        |-- core.StatProcessor(cpu.NewProcessor(filename), chan)
        |-- core.StatProcessor(ram.NewProcessor(filename), chan)
        |-- core.StatProcessor(dsk.NewProcessor(filename), chan)
        |            // disk io?, disk storage?
        |-- core.StatProcessor(net.NewProcessor(filename), chan)
        |            // does this use the /proc fs?
    +-----------+
    | processor |- StatProcessor(statProcessor, chan)
    +-----------+
        |-- Process() -> []core.Stat{}
    +-----------------+
    | ResettingReader | Open(filename)
    +-----------------+ 
        |-- Read() -> [][]string
    +---------+
    | io.File | Open(filename)
    +---------+
        |-- Seek()
        \-- Read()

