# Credit Card Type Detector

A Go implementation of the CS50 credit card validation problem that detects and validates AMEX, MasterCard, and Visa credit cards using the Luhn algorithm.

## Features

- **AMEX Detection**: 15-digit cards starting with 34 or 37
- **MasterCard Detection**: 16-digit cards starting with 51-55
- **Visa Detection**: 13 or 16-digit cards starting with 4
- **Luhn Algorithm Validation**: All cards must pass checksum validation
- **Comprehensive Error Handling**: Invalid cards return "INVALID"

## Usage

```bash
# Run the program
go run .

# Run tests
go test -v

# Check test coverage
go test -cover
```

## Test Coverage

The test suite includes comprehensive coverage (94.1%) with the following test categories:

### DetectCardType Tests (27 test cases)
- **Valid Cards**: AMEX, MasterCard, and Visa with correct formats and checksums
- **Invalid Cards**: Wrong lengths, invalid checksums, incorrect identifying digits
- **Edge Cases**: Zero, single digits, boundary conditions

### Helper Function Tests

#### TrimNumber Tests (7 test cases)
- Extracting first N digits from numbers
- Edge cases: zero digits, more digits than available

#### NumDigits Tests (7 test cases)
- Counting digits in numbers from single digits to 16-digit cards
- Special case: zero returns 0 digits

#### IsValidLuhn Tests (13 test cases)
- Valid credit card numbers that pass Luhn algorithm
- Invalid numbers with wrong checksums
- Simple test cases for algorithm verification

## Test Results

All tests based on CS50 problem requirements:
- ✅ Identifies 378282246310005 as AMEX
- ✅ Identifies 371449635398431 as AMEX  
- ✅ Identifies 5555555555554444 as MASTERCARD
- ✅ Identifies 5105105105105100 as MASTERCARD
- ✅ Identifies 4111111111111111 as VISA
- ✅ Identifies 4012888888881881 as VISA
- ✅ Identifies 4222222222222 as VISA
- ✅ Properly rejects invalid cards with wrong lengths, checksums, or identifying digits

## Fixed Issues

During development, two critical bugs were identified and fixed:
1. **MasterCard Logic**: Changed `||` to `&&` in range check (51-55)
2. **Visa Detection**: Changed from checking first 2 digits to first 1 digit

## Example Output

```
Number: 4111111111111111
VISA

Number: 378282246310005  
AMEX

Number: 1234567890
INVALID
```