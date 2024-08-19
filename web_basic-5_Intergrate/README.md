## 架構


* Model Layer: 
    
    用於定義**資料格式**與屬性，此階層所定義的物件會在其他三層都會用到。
* Repository Layer: 
    
    用於與**外部 service 互動**的 adapter(外部service 像是DB，是此server使用的服務，並非使用者喔)。，例如要與 mariadb、mongo 或是 redis 等 data storage 進行互動，就將方法封裝在此階層內。
* Service Layer:

    用於封裝商業邏輯方法，在**系統內**有用到的邏輯判斷或是演算法都會放在此階層內。
* Delivery Layer:

    用於封裝接口給**外部使用者**使用的方法，例如要提供 rest、gRPC 或是 command line 的方法都會封裝在此階層內。

### 此系統架構
```
/
    /model (model layer)
        - model.go
    /module
        /user
            /delivery (給外部使用者)
                - handler.go
            / repository
                - repository.go (實作)
            / service
                - service.go
                
            - repository.go (和 DB 互動的interface)
            - service.go 

```