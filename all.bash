#!/bin/bash

set -e

platform='unknown'
unamestr=`uname`
if [[ "$unamestr" == 'Linux' ]]; then
   platform='linux'
elif [[ "$unamestr" == 'MINGW32_NT-6.2' ]]; then
   platform='windows'
elif [[ "$unamestr" == 'MINGW64_NT-10.0' ]]; then
   platform='windows'
elif [[ "$unamestr" == 'Darwin' ]]; then
   platform='darwin'
fi

echo "platform = $platform"

install_and_test(){
    HERE=`pwd`
    PKG=$1
    DOTEST=$2
    echo
    echo
    echo "[1;32m>>> compiling $PKG <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<[0m"
    cd $PKG
    touch *.go
    go install
    if [ "$DOTEST" -eq 1 ]; then
        go test
    fi
    cd $HERE
}

for p in chk io utl plt; do
    install_and_test $p 1
done

if [[ $platform == 'linux' ]]; then
    install_and_test mpi 0
fi

for p in la fdm num fun gm graph ode opt rnd tsr; do
    install_and_test $p 1
done

if [[ $platform == 'linux' ]]; then
    install_and_test vtk 0
fi
