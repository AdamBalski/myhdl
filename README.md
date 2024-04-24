# MyHDL
a hardware description language

## It is often beneficial to discuss the pros and cons of technologies to make informed decisions and choose wisely
Let's compare MyHDL against one of the most used HDL languages, Verilog.

Pros of MyHDL over Verilog:
* it is/will be/was really fun building this suite

Pros of Verilog over MyHDL:
* I can't seem to find one...

## MyHDL file description:
* one gate/chip per file,
* by convention the filename ends with '.mhf' as an abbreviation of MyHdl File
* all white spaces outside of quotation marks get ommitted
* a comment starts with a '//' and lasts until the end of line
* `name` property describes the chip's name (there has to be exactly one per file)
    * example: `name Not`
* `uses` property describes a chip used within the chip we're describing
    * it describes the type of chip used and space-delimited identifiers which are going to be used within the file
    * example: `uses Nand nand`
* `inputs` or `input` property is a space-delimited list of inputs' names
    * example: `input in`
* `outputs` property is a space-delimited list of outputs' names
    * example: `output out`
* `propagation` property describes the way in which the signals propagate within the chip
    * example: 
    ```
    nand.inputA :- in
    nand.inputB :- in
    out :- nand.output
    ```
* `deterministic` and `nondeterministic`
    * an example of a deterministic gate would be an OR gate
    * an example of a non deterministic gate is a latch
    * by default a chip is deterministic if and only if there exists a topological sorting of signalpropagation, including within the sub-chips
    * this property decides whether the chip can be compiled to the 'cmh' format
    * a chip can't have `deterministic` and `nondeterministic` set at the same time`

## Example MyHDL files:
### 'and.mhf'
```
name And
deterministic
uses Not not
uses Nand nand
inputs inA, inB
output out
nand.inA :- inA
nand.inB :- inB
not.in :- nand.out
out :- not.out
```
### 'or.mhf'
```
name Or
deterministic
inputs inA, inB
output out
uses Not notOnInputA notOnInputB
notOnInputA.in :- inA
notOnInputB.in :- inB
uses And and
and.inA :- notOnInputA.out
and.inB :- notOnInputB.out
uses Not notAfterAnd
notAfterAnd.in :- and.out
out :- notAfterAnd.out
```
### 'sr_latch.mhf'
```
// Nand SrLatch
name SrLatch
nondeterministic
inputs set, reset
output state
uses Nand setNand, resetNand
resetNand.inA :- reset
resetNand.inB :- setNand.out
setNand.inA :- resetNand.out
setNand.inB :- set
state :- resetNand.out
```

## Compiled MyHDL (the aforementioned 'cmh' format)
A chip can be compiled to cmh if and only if it is deterministic

Note: In order for a chip to be deterministic all its sub-chips have to be as well

The format:
* first line is the name of the chip, the number of inputs, then the number of outputs, and then all inputs and outputs' names
* second line:
    * let 'n' be the number of inputs and 'm' be the number of outputs
    * `2^n` space delimited values describing outputs for each possible input state
    * let `output` be the number at the `input`'st position (zero-indexed)
    * `(input >> i) & 1` is the state of the `i`'th input
    * `(output >> o) & 1` is the state of the `o`'th output

Example describing a NAND gate:
```
Nand 2 1 inA inB out
1 1 1 0
```
Explanation:
| position | position in binary | inB | inA | out|
|:--------:|:------------------:|:---:|:---:|:--:|
| 0 | 00 | 0 | 0 | 1 |
| 1 | 01 | 0 | 1 | 1 |
| 2 | 10 | 1 | 0 | 1 |
| 3 | 11 | 1 | 1 | 0 |
    
