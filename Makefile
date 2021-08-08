CXX=g++
INCLUDES=-I. -I../src/ -I../
CXXFLAGS=-fPIC -pthread -Wall -std=c++0x -std=c++11 -O2 -march=x86-64 $(INCLUDES)
LDFLAGS=-shared
OBJS=sentencepiece_wrapper.o
TARGET=libsentencepiece.so
SRC=src/CMakeFiles/sentencepiece.dir

all: $(TARGET)

$(OBJS): sentencepiece_wrapper.h sentencepiece_wrapper.cc
	$(CXX) $(CXXFLAGS) -c sentencepiece_wrapper.cc

$(TARGET): $(OBJS)
	$(CXX) $(LDFLAGS) -o $(TARGET) $(OBJS) src/libsentencepiece.a
