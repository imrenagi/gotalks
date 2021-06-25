
## Notes

## Speed

## Watching for "sleeping nails"

<- time.After(10 * time.Second)

## Use test doubles if test rely on expensive resources

## Using tag for running excluding integration test

don't

should use t.Skip()
* so that people might notice a test was being skipped 

## Avoid Duplication

CleanUp() function in go

## Iteration on test might be hard to maintain

You might need to check whether there is a particular number from a slice returned by a func

.code testing/simple/contains.go /STARTCONTAINS10/,/STOPCONTAINS10/ 

## Use assertion framework

Utilize assertion like `testify` for checking the oracle

.code testing/simple/contains.go /START/,/STOP/