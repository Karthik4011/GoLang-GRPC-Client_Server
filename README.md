Project Description : https://www.csee.umbc.edu/~kalpakis/courses/621-sp22/project/GoTokens-1.pdf

Steps:
1) We must first generate a proto file, which we must then execute with server_script.sh

2) We must develop client server go files after creating the proto file and implement the functionality based on the proto file in those go files.

3) Execute the server_script.sh file first, then the client_script.sh file.


libraries used are:
"context"
"crypto/sha256"
"encoding/binary"
"fmt"
"log"
"net"
"strconv"
"strings"
"sync"
"os"
"time"


NOTE: server_script.sh and client_script.sh are the two bash srcripts used to execute the project

Github Repo Link: https://github.com/Karthik4011/GoLang-GRPC-Client_Server


References used :
https://grpc.io/docs/languages/go/basics/
https://developers.google.com/protocol-buffers/docs/gotutorial