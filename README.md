# CodeGen
A code generator program built by golang.

## IDEA

The idea came from kubectl in Kubernetes. When creating a pod or other object by
kuebctl, we use:

    kubectl create -f test.yaml

Kubernetes let user use config files in yaml or json format to 
specify their own configuration, which is very flexible and 
user-friendly. 

So I came up with an idea if I could create a binary to generate Java code in the same way.


## Examples

An example test.yaml file: 

    name: Test
    type:  Controller
    ctype: class
    mode:  public
    methods:
      - mname: show
        mmode:  private static
        mparameters:
          - pname: age
            ptype: int
          - pname: name
            ptype: String
        mreturn:
          pname: user
          ptype: String
    
    variables:
      - vname:  name
        vtype:  String
        vmode:  public
        vvalue: puyangsky


Use

    go run Main.go -t yaml -o ./ -f test.yaml
    
Output:
    
    [INFO] FileName: test.yaml
    [INFO] Output: ./
    [INFO] GenType: yaml
    -----
    [DEBUG] Config content:
    ------
    name: Test
    type:  Controller
    ctype: class
    mode:  public
    methods:
      - mname: show
        mmode:  private static
        mparameters:
          - pname: age
            ptype: int
          - pname: name
            ptype: String
        mreturn:
          pname: user
          ptype: String
    
    variables:
      - vname:  name
        vtype:  String
        vmode:  public
        vvalue: puyangsky
    -----
    [DEBUG] Class content:
    ------
    ClassName: Test
    Type: Controller
    ClassType: class
    AccessMode: public
    Methods len: 1
    Variables len: 1