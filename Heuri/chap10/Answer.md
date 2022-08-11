## Q1 

switch, case, case, default

## Q2

긍정적인 평가

## Q3

```go
func GetDirection(angle float64) Direction {
  switch {
    case angle >= 315, angle >=0 && angle < 45:
  		return North
    case angle >= 45 && angle < 135:
    	return East
    case angle >= 135 && angle < 225:
    	return South
    case angle >= 225 && angle <= 315:
    	return West
    default:
    	return None
  } 
  	
  case :
}
```

