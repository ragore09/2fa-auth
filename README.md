# 🔐2FA Algorithm Implementation
Implementation of the 2FA Algorithm. 

## Getting started
Add the provider and secret entries in the `secrets` file.  
Your file must look like this:
```
Provider AAAABBBBCCCCDDD
AnotherProvider AAAABBBBCCCCDDD
```

Once this is done, run 
```
go run main.go utils.go
```

This will generate the following output:
```
Provider: 194872
AnotherProvider: 120931
```

You're ready to go, use the generated One-time Password.
