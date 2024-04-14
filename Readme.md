### integer data type
- int8 and uint8 -> unsigned int, start from 0
- int16 and uint16
- int32 and uint32
- int64 and uint64
### float data type
- float32
- float64
- complex64
- complex128
### alias for data type
- byte -> uint8
- rune -> int32
- int Minimal int32
- uint -> Minimal uint32
### boolean data type
- true or false
### string data type
- declaration -> `string, "string"`
### function for string
- len("string") -> character length in string
- "string"[number] -> get character by index, and return byte of strings
### variable
- variable is to storing data
- how to write variable -> `var name`
- multiple declaration variable -> `var (name = "yourName", age = 24)`
### constant
- is variable constant, can't to re assign
- create constant must assign the value
### data type conversion
- example want to conversion from int32 -> int64
### type declaration
- the ability to create new object with same type
- the purpose is to understanding object by create alias
### create array
- syntax `array := [...]string{"satu", "dua", "tiga"}`
- get array length using `len(array)`
### slice
- we can slice from array using `slice := array[start:end]`, if start or end is not init it will start from 0 and end until last index of array
- init slice using make => `slice_name := make([]type, length, capacity)`
- append element to slice
  - `slice_name = append(slice_name, element1, element2, ...)`
  - `slice3 = append(slice1, slice2...)`
- we can copy slice using `copy(dest, src)`

### map

- map is like object in javascript

- create map using syntax

  ```go
  variable := map[typeKey]typeValue{
      key: value
  }
  ```

- access map => `variable[key]`

- change value or add value => `variable[key] = value`

- length of map => `len(variable)`

- create blank map using make function

  ```go
  mapName := make(map[typeKey]typeValue)
  ```

delete map using key => `delete(mapName, key)`