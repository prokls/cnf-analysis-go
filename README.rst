cnf-analysis-go
===============

:author:     Lukas Prokop
:date:       April to August 2016
:version:    1.0.0
:license:    CC-0

cnf-analysis-go is a tool to analyze DIMACS CNF files.
Those files are commonly used to decode SAT problems and
a few basic features of a CNF file might tell you something
about the problem stated.

This tool evaluates a list of features which are thoroughly
described by the project and stores them in a JSON file.

How To Use
----------

To install use ` ``go get`` and ``go install`` <https://golang.org/cmd/go/>`_::

    $ go get github.com/prokls/cnf-analysis-go
    $ go install github.com/prokls/cnf-analysis-go
    $ ls -hl $GOROOT/bin/cnf-analysis-go
    -rwxrwxr-x 1 prokls prokls 3.4M Aug 14 16:18 /opt/go-1.6.2/bin/cnf-analysis-go

Then the command line tool to analyze CNF files is available::

    $ echo "p cnf 3 2\n1 -3 0\n2 3 -1 0\n" > example.cnf
    $ cnf-analysis-go example.cnf
    2016/08/14 14:45:42 considering example.cnf
    2016/08/14 14:45:42 writing file example.stats.json
    $ cat example.stats.json
    [{"@cnfhash":"cnf2$7d16f8d71b7097a2f931936ae6d03d738117b2c6","@filename":"exa
    $ cat example.stats.json | jq
    [
      {
        "@cnfhash": "cnf2$7d16f8d71b7097a2f931936ae6d03d738117b2c6",
        "@filename": "example.cnf",
        "@md5sum": "04f6bf2c537242f15082867e66847bd7",
        "@sha1sum": "23dd9e64ae0fb4806661b49a31e7f5e692f2d045",
        "@timestamp": "2016-08-14T12:45:42.262244157",
        "@version": "1.0.0",
        "featuring": {
          "clause_variables_sd_mean": 0.9082483053207397,
          "clauses_count": 2,
          "clauses_length_largest": 3,
          "clauses_length_mean": 2.5,
          "clauses_length_median": 2.5,
          "clauses_length_sd": 0.5000000000000001,
          "clauses_length_smallest": 2,
          "connected_literal_components_count": 3,
          "connected_variable_components_count": 1,
          "definite_clauses_count": 1,
          "existential_literals_count": 1,
          "existential_positive_literals_count": 1,
          "false_trivial": true,
          "goal_clauses_count": 0,
          "literals_count": 5,
          "literals_frequency_0_to_5": 1,
          "literals_frequency_5_to_10": 0,
          "literals_frequency_10_to_15": 0,
          "literals_frequency_15_to_20": 0,
          "literals_frequency_20_to_25": 0,
          "literals_frequency_25_to_30": 0,
          "literals_frequency_30_to_35": 0,
          "literals_frequency_35_to_40": 0,
          "literals_frequency_40_to_45": 0,
          "literals_frequency_45_to_50": 0,
          "literals_frequency_50_to_55": 5,
          "literals_frequency_55_to_60": 0,
          "literals_frequency_60_to_65": 0,
          "literals_frequency_65_to_70": 0,
          "literals_frequency_70_to_75": 0,
          "literals_frequency_75_to_80": 0,
          "literals_frequency_80_to_85": 0,
          "literals_frequency_85_to_90": 0,
          "literals_frequency_90_to_95": 0,
          "literals_frequency_95_to_100": 0,
          "literals_frequency_entropy": 2.5,
          "literals_frequency_largest": 0.5,
          "literals_frequency_mean": 0.4166666666666667,
          "literals_frequency_median": 0.5,
          "literals_frequency_sd": 0.18633899812498247,
          "literals_frequency_smallest": 0,
          "literals_occurence_one_count": 5,
          "nbclauses": 2,
          "nbvars": 3,
          "negative_literals_in_clause_largest": 1,
          "negative_literals_in_clause_mean": 1,
          "negative_literals_in_clause_smallest": 1,
          "negative_unit_clause_count": 0,
          "positive_literals_count": 3,
          "positive_literals_in_clause_largest": 2,
          "positive_literals_in_clause_mean": 1.5,
          "positive_literals_in_clause_median": 1.5,
          "positive_literals_in_clause_sd": 0.5000000000000001,
          "positive_literals_in_clause_smallest": 1,
          "positive_negative_literals_in_clause_ratio_entropy": 0.8899749834391559,
          "positive_negative_literals_in_clause_ratio_stdev": 0.08333334326744612,
          "positive_negative_literals_in_clause_ratio_mean": 0.5833333730697632,
          "positive_unit_clause_count": 0,
          "tautological_literals_count": 0,
          "true_trivial": true,
          "two_literals_clause_count": 1,
          "variables_frequency_0_to_5": 0,
          "variables_frequency_5_to_10": 0,
          "variables_frequency_10_to_15": 0,
          "variables_frequency_15_to_20": 0,
          "variables_frequency_20_to_25": 0,
          "variables_frequency_25_to_30": 0,
          "variables_frequency_30_to_35": 0,
          "variables_frequency_35_to_40": 0,
          "variables_frequency_40_to_45": 0,
          "variables_frequency_45_to_50": 0,
          "variables_frequency_50_to_55": 1,
          "variables_frequency_55_to_60": 0,
          "variables_frequency_60_to_65": 0,
          "variables_frequency_65_to_70": 0,
          "variables_frequency_70_to_75": 0,
          "variables_frequency_75_to_80": 0,
          "variables_frequency_80_to_85": 0,
          "variables_frequency_85_to_90": 0,
          "variables_frequency_90_to_95": 0,
          "variables_frequency_95_to_100": 2,
          "variables_frequency_entropy": 0.5,
          "variables_frequency_largest": 1,
          "variables_frequency_mean": 0.8333333333333334,
          "variables_frequency_median": 1,
          "variables_frequency_sd": 0.23570226039551584,
          "variables_frequency_smallest": 0.5,
          "variables_largest": 3,
          "variables_smallest": 1,
          "variables_used_count": 3
        }
      }
    ]


Performance
-----------

``cnf-analysis-go`` creates a few workers (its number can be
set by ``-u``) and puts CNF files in a queue. Whenever a worker
has finished its work, it receives the next CNF file. A worker
is modelled by a goroutine. Per default 4 workers are used.

This implementation is recognizably faster than ``cnf-analysis-go``
and uses less memory. The following data was retrieved for
SAT competition 2016 files:

* ``esawn_uw3.debugged.cnf`` (1.4 GB) in *app16* took 1 hour and 13 minutes
* ``bench_573.smt2.cnf`` (1.6 MB) in *Agile* took 1 second

I am using my Thinkpad x220t with 16GB RAM and an Intel Core
i5-2520M CPU (2.50GHz) as reference system here.

Memory
------

The major advantage of ``cnf-analysis-go`` over ``cnf-analysis-py``
is its memory usage.

* ``esawn_uw3.debugged.cnf`` (1.4 GB) in *app16* used approx. 3 GB RAM
* ``bench_573.smt2.cnf`` (1.6 MB) in *Agile* uses approx. 2 MB memory

Considering a factor of 10 for ``cnf-analysis-py``, I assume a factor
(ratio of filesize to RAM usage) of at most 3 for ``cnf-analysis-go``.
Its memory usage is generally much more constant, but on the other hand
the results are less accurate.

Because of these data I recommend up to 6 units for 8 GB RAM or
12 units for 16 GB RAM.

I achieved my goal to make this implementation memory-efficient.

Dependencies
------------

* `golang <http://golang.org/>`_ or specifically 'go1.6.2 linux/amd64' was used

It has only one external dependency, namely

* `cnf-hash-go <https://github.com/prokls/cnf-hash-go/>`_

Command line options
--------------------

``--ignore c --ignore x``
  Ignore any lines starting with "c" or "x".
  If none is specified "c" and "%" is ignored.
``--units 4`` or ``-u 4``
  use (at most) 4 parallel units
``--no-hashes`` or ``-n``
  skip hash computations (SHA1, MD5, cnfhash)
``--fullpath`` or ``-p``
  print full path, not basename
``--skip-existing`` or ``-s``
  skip stats computation if stats.json already exists

DIMACS files
------------

DIMACS files are read by skipping any lines starting with characters
from ``--ignore``. The remaining content is parsed (header line with
``nbvars`` and ``nbclauses``) and in the remaining line, integers are
retrieved and passed over. Hence the parser yields a sequence of
literals.

Features
--------

Features are documented in my paper "Analyzing CNF benchmarks".

Cheers,
prokls
