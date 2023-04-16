# What is a variable
    it is a container for storing data.
### What does it look like
    py:  [name] = [value] # a = 1
    js:  [let/var] [name] = [value]; // let a = 1;
    Go: var [name] [type] = [value] // var a int32 = 1
### What data can be stored
    Integers: Whole numbers without decimal places
        ETC: 1, 2, skip a few, 42_108_642_164_021
        101 = 5
        [128, 64, 32, 16, 8, 4, 2, 1]
        1
        10
        11
        100
        101
        110
        111
        1000
    
    Floating point numbers: Numbers that can contain decimal places
        ETC: 1.0, 1.5, 506_285.156_028
    
    Char: a number that corresponds to ascii text. This is represented as an eight bit number in memory.
        ETC: 'a', 'B', 'Æ', ' ', '\n'

    String: a series of Characters that form a sequence.
        ETC: 'U R Bad'

    Boolean: a representation containing two values: True or False
        ETC: True, False

    Complex variables: Arrays, structures, *objects, *functions... basically anything.
        * most languages

### Scope

    variables live for the duration of the code block they live in.

    This is best taught through examples.

    EX: Language Python.
    var1 = 'scope 1'
    if True:
        var2 = 'scope 2'
        print(var1) # scope 1
        print(var2) # scope 2
    print(var1) # scope 1
    print(var2) # ERROR Variable does not exist.

    This concept exists for all major language with some minor differences between them.