PROG_NAME=c-cplex-example-executable
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
CFLAGS= $(INC)
COPTFLAGS= -m64 -O3 -fPIC -fexceptions -DNDEBUG -DIL_STD -DLONG_MAX=0x7FFFFFFFL
CPLEXINCDIR   = -I$(CPLEXDIR)/cplex/include -I$(CPLEXDIR)/concert/include
CPLEXFLAGS  = $(CFLAGS) $(COPTFLAGS) $(CPLEXINCDIR) 
CC = gcc

MODEL_H=$(SRC_DIR)/model.h
MODEL_SRC=$(SRC_DIR)/model.c
MODEL_OBJ=$(PROG_BIN)/model.o

model: $(MODEL_SRC) $(MODEL_H)
	$(CC) $(CPLEXFLAGS) -c $(MODEL_SRC) -o $(MODEL_OBJ)

main: 
	$(CC) $(CPLEXFLAGS) -c $(MAIN_SRC) -o $(PROG_BIN)/main.o

clean_bin:
	rm -f $(PROG_BIN)/*.o $(PROG_BIN)/*.a
clean:
	rm -f Makefile.deps $(PROG_NAME).seq $(PROG_BIN)/*.o $(PROG_NAME)
c-cplex-lib: model
	libtool -static -o $(PROG_BIN)/libmodel.a $(PROG_BIN)/model.o

all: c-cplex-lib clean_bin

Makefile.deps:
	$(CC) $(CPLEXFLAGS) -MM $(SRC_DIR)/*.c > Makefile.deps
-include Makefile.dep