package main

type RespType byte

const (
	SimpleString   RespType = '+'
	Error          RespType = '-'
	Integer        RespType = ':'
	BulkString     RespType = '$'
	Array          RespType = '*'
	Null           RespType = '_'
	Boolean        RespType = '#'
	Double         RespType = ','
	BigNumber      RespType = '('
	BulkError      RespType = '!'
	VerbatimString RespType = '='
	Map            RespType = '%'
	Attribute      RespType = '|'
	Set            RespType = '~'
	Push           RespType = '>'
)
