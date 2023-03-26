# simulatooor
This repo is for demo simulations

The script basically takes all the list of variables and contract address which need to be indexed.

To run the script, you need to mention all the variables in monitor.json

The querioor supports the datatypes such as int, bool, string, address, map, nested maps, array etc.


### Setup
```
{
    "contractAddress" : "0x3126D03e98bb95a7d4046bA8A64369E6656Fe448",
    "no_of_variables" : 10,
    "variables" : [
        {"name" : "val1"},
        {"name" : "bool1"},
        {"name" : "str1"},
        {"name" : "addr1"},
        {"name" : "arr", "key": "0"},
        {"name" : "arr", "key": "1"},
        {"name" : "arr", "key": "2"},
        {"name" : "arr", "key": "3"},
        {"name" : "arr", "key": "4"},
        {"name" : "balances", "key": "0x0f3aac271357DdE397c6a59204Cf5FD2Ac7f78ea"},
        {"name" : "balances2", "key": "0x0f3aac271357DdE397c6a59204Cf5FD2Ac7f78ea","deep_key": "0x0f3aac271357DdE397c6a59204Cf5FD2Ac7f78ea"}
    ]
}
```

In the main.go file, inside the ```SendPostRequest``` function, update the URL with your queriooor service URL.

### Run
To run the service, 

``` go run main.go ```

### Output
![WhatsApp Image 2023-03-26 at 8 41 35 PM](https://user-images.githubusercontent.com/32795247/227785715-fa0abe43-29d6-45b2-b0c4-7d5e6819d52c.jpeg)




