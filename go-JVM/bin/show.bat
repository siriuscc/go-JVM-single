




@choice /m "1, compile go-JVM?"

@if %errorlevel%==2 goto compileJava

:compile

@G:
cd G:\graduationDesign\go-JVM\src
@set GOPATH=G:\graduationDesign\go-JVM
@set GOBIN=%GOPATH%/bin
go install
cd G:\graduationDesign\go-JVM/bin
@del go-JVM.exe
@rename src.exe go-JVM.exe


:compileJava

@choice /m "2, compile java file to class file?"
@if %errorlevel%==2 goto test

cd G:\graduationDesign\go-JVM\bin/java

@for /f "delims=" %%i in ('dir /b *.java') do (

    javac -encoding utf8 -d ../classes %%i
)




:test

@choice /m "3, test run class file with go-JVM?"
@if %errorlevel%==2 goto end

cd G:\graduationDesign\go-JVM\bin


@echo Array test:
go-JVM -cp=classes ArrayDemo
java   -cp classes ArrayDemo


@echo "Auto box test:"
go-JVM -cp=classes BoxTest
java   -cp classes BoxTest

@echo "Bubble sort test:"
go-JVM -cp=classes BubbleSortTest
java   -cp classes BubbleSortTest



@echo "Class test:"
go-JVM -cp=classes ClassDemo
java   -cp classes ClassDemo

@echo "Clone test:"
go-JVM -cp=classes CloneTest
java   -cp classes CloneTest


@echo "Exception catch test:"
go-JVM -cp=classes ExceptionCatch
java   -cp classes ExceptionCatch


@echo "No main method test:"
go-JVM -cp=classes ExceptionDemo
java   -cp classes ExceptionDemo



@echo "Fibonacci test:"
go-JVM -cp=classes FibonacciTest
java   -cp classes FibonacciTest


@echo "Get Class reflect test:"
go-JVM -cp=classes GetClassTest
java   -cp classes GetClassTest


@echo "Get HashCode test:"
go-JVM -cp=classes HashCodeDemo
java   -cp classes HashCodeDemo


@echo "Hello World:"
go-JVM -cp=classes HelloWorld
java   -cp classes HelloWorld


@echo "Invoke method test:"
go-JVM -cp=classes InvokeDemo
java   -cp classes InvokeDemo


@echo "instance and instanceof test:"
go-JVM -cp=classes MyObject
java   -cp classes MyObject


@echo "Parse args[0] to int test:"
go-JVM -cp=classes ParseIntTest
java   -cp classes ParseIntTest


@echo "Print args test:"
go-JVM -cp=classes PrintArgs
go-JVM -cp=classes PrintArgs abc jkl
java -cp classes PrintArgs
java -cp classes PrintArgs abc jkl


@echo "String Reflect test:"
go-JVM -cp=classes StringReflect
java   -cp classes StringReflect


@echo "String Reflect test:"
go-JVM -cp=classes StringTest
java   -cp classes StringTest


:end


