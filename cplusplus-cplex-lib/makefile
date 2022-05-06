PROG_NAME=cplusplus-cplex-example-executable
PROG_BIN=bin
SRC_DIR=src
SYSTEM = x86-64_osx
LIBFORMAT = static_pic
CPLEXDIR      = /Applications/CPLEX_Studio1210/
CONCERTDIR    = /Applications/CPLEX_Studio1210/concert
CPLEXLIBDIR   = $(CPLEXDIR)/cplex/lib/$(SYSTEM)/$(LIBFORMAT)
CONCERTLIBDIR = $(CONCERTDIR)/lib/$(SYSTEM)/$(LIBFORMAT)
CLNFLAGS  = -L$(CPLEXLIBDIR) -L$(CONCERTLIBDIR) -lilocplex -lconcert -lcplex -m64 -lm -lpthread
INC= -I$(SRC_DIR)
CPPFLAGS= -std=c++17 $(INC)
COPTFLAGS= -m64 -O3 -fPIC -fexceptions -DNDEBUG -DIL_STD -DLONG_MAX=0x7FFFFFFFL
CPLEXINCDIR   = -I$(CPLEXDIR)/cplex/include -I$(CPLEXDIR)/concert/include
CPLEXFLAGS  = $(CPPFLAGS) $(COPTFLAGS) $(CPLEXINCDIR) 
CC = gcc
MAIN_SRC=$(SRC_DIR)/main.cpp

MODEL_H=$(SRC_DIR)/model.h
MODEL_SRC=$(SRC_DIR)/model.cpp
MODEL_OBJ=$(PROG_BIN)/model.o

model: $(MODEL_SRC) $(MODEL_H)
	$(CC) $(CPLEXFLAGS) -c $(MODEL_SRC) -o $(MODEL_OBJ)

clean_bin:
	rm -f $(PROG_BIN)/*.o $(PROG_BIN)/*.a
clean:
	rm -f Makefile.deps $(PROG_NAME).seq $(PROG_BIN)/*.o $(PROG_NAME)
all: cplusplus-cplex-lib clean_bin
cplusplus-cplex-lib: model
	libtool -static -o $(PROG_BIN)/libmodelcplusplus.a $(PROG_BIN)/model.o

Makefile.deps:
	$(CC) $(CPLEXFLAGS) -MM $(SRC_DIR)/*.cpp > Makefile.deps
-include Makefile.deps
