package benchvm

import (
	"log"

	"github.com/benchlab/benchvm/bvmstate"
	"github.com/benchlab/bvmCreate"
)

const bvm = `
name = "BenchVM"
author = "BlockStarters Lab"
description = "BenchChain's Core Virtual Machine Engine For Smart Contracts."

category("MATH"){

    description = "Basic "

    spec("ADD", "01"){
        description = "Addition"
        validate = 2
        oxygen = 10
    }

    spec("SUB", "02"){
        description = "Subtraction"
        validate = 2
        oxygen = 10
    }

    spec("MUL", "03"){
        description = "Multiplication"
        validate = 2
        oxygen = 10
    }

    spec("DIV", "04"){
        description = "Multiplication"
        validate = 2
        oxygen = 10
    }

    spec("MOD", "05"){
        validate = 2
    }

    spec("ADDMOD", "06"){
        validate = 3
    }

    spec("MULMOD", "07"){
        validate = 3
    }

    spec("CONCAT", "0A"){
        description = "Concatenates two byte arrays"
        validate = 2
    }
}

category("CRYPTO LIBRARY"){

    description = "Cryptographic functions."

    spec("SHA3", "30"){

    }
}

category("MATH COMPARE"){

    description = "Comparison operations."

    spec("LT"){

    }

    spec("GT"){

    }

    spec("SLT"){

    }

    spec("SGT"){

    }

    spec("EQ"){

    }

    spec("ISZERO"){

    }
}

category("CHAIN LOGIC"){

    spec("AND"){

    }

    spec("OR"){

    }

    spec("XOR"){

    }

    spec("NOT"){

    }

}

category("CONFIG"){

    spec("ADDRESS"){

    }

    spec("BALANCE"){

    }

    spec("ORIGIN"){

    }

    spec("CALLER"){

    }

    spec("CALLVALUE"){

    }

    spec("CALLDATALOAD"){

    }

    spec("CALLDATASIZE"){

    }

    spec("CALLDATACOPY"){

    }

    spec("CODESIZE"){

    }

    spec("CODECOPY"){

    }

    spec("FUELCOST"){

    }

    spec("EXTCODESIZE"){

    }

    spec("EXTCODECOPY"){

    }

}

category("CHAIN BLOCK"){

    description = "Operations which relate to the current block."

    spec("BLOCKHASH"){

    }

    spec("COINBASE"){

    }

    spec("TIMESTAMP"){

    }

    spec("NUMBER"){

    }

    spec("DIFFICULTY"){

    }

    spec("OXYGENLIMIT"){

    }

}

category("MACHINE FLOW"){

    spec("BOUNCE", "70"){

    }

    spec("BOUNCEI", "71"){

    }

    spec("COMP", "72"){

    }

    spec("OXYGEN", "73"){

    }

    spec("BOUNCEDEST", "74"){

    }

    spec("STOP", "75"){

    }

    spec("RETURN", "76"){

    }
}

category("MACHINE MEMORY"){
    spec("GET", "40"){

    }

    spec("SET", "41"){

    }

    spec("STORE", "42"){

    }

    spec("LOAD", "43"){

    }
    spec("MSIZE", "44"){

    }
}

category("BVM"){
    spec("POP", "60"){

    }
    spec("PUSH", "61"){

    }
    spec("DUP", "62"){

    }
    spec("SWAP", "63"){

    }
}

category("MACHINE STRUCTURES"){

    description = "Creating structures."

    spec("ARRAY", "A0"){

    }

    spec("MAP", "A1"){

    }
}
`

var (
	oxygens    = map[string]bvmCreate.FuelFunction{}
	deploys = map[string]bvmCreate.ExecuteFunction{
		// chainFormulas functions
		"ADD":    bvmstate.Add,
		"SUB":    bvmstate.Sub,
		"DIV":    bvmstate.Div,
		"MUL":    bvmstate.Mul,
		"MOD":    bvmstate.Mod,
		"ADDMOD": bvmstate.AddMod,
		"MULMOD": bvmstate.MulMod,
		"CONCAT": bvmstate.Concat,
		// mathComparison functions
		"LT":     bvmstate.Lt,
		"GT":     bvmstate.Gt,
		"SLT":    bvmstate.SLt,
		"SGT":    bvmstate.SGt,
		"EQ":     bvmstate.Eq,
		"ISZERO": bvmstate.IsZero,
		// chainLogic functions
		"AND": bvmstate.And,
		"OR":  bvmstate.Or,
		"NOT": bvmstate.Not,
		"XOR": bvmstate.Xor,
		// chainBlock functions
		"HASH":       bvmstate.blockHash,
		"COINBASE":   bvmstate.Coinbase,
		"TIMESTAMP":  bvmstate.Timestamp,
		"NUMBER":     bvmstate.Number,
		"DIFFICULTY": bvmstate.Difficulty,
		"OXYGENLIMIT":  bvmstate.OxygenLimit,
		// config functions
		"ADDRESS":      bvmstate.Address,
		"ORIGIN":       bvmstate.Origin,
		"BALANCE":      bvmstate.Balance,
		"CALLER":       bvmstate.Caller,
		"CALLVALUE":    bvmstate.CallValue,
		"CALLDATALOAD": bvmstate.CallDataLoad,
		"CALLDATASIZE": bvmstate.CallDataSize,
		"CALLDATACOPY": bvmstate.CallDataCopy,
		"CODESIZE":     bvmstate.CodeSize,
		"CODECOPY":     bvmstate.CodeCopy,
		"OXYGENPRICE":    bvmstate.OxygenPrice,
		"EXTCODESIZE":  bvmstate.ExtCodeSize,
		"EXTCODECOPY":  bvmstate.ExtCodeCopy,
		// cryptoLib functions
		"SHA3": bvmstate.S3,
		//machine functions
		"POP.": bvmstate.Pop,
		"PUSH": bvmstate.Push,
		"DUP":  bvmstate.Dup,
		"SWAP": bvmstate.Swap,
		// external functions
		"LOG": bvmstate.Log,
		// machineFlow functions
		"BOUNCE":     bvmstate.Bounce,
		"BOUNCEI":    bvmstate.BounceI,
		"COMP":       bvmstate.Comp,
		"OXYGEN":     bvmstate.Oxygen,
		"MEMSIZE":  bvmstate.MemSize,
		"BOUNCEDEST": bvmstate.BounceDest,
		// machineMemory functions
		"SET":   bvmstate.Set,
		"GET":   bvmstate.Get,
		"LOAD":  bvmstate.Load,
		"STORE": bvmstate.Store,
	}
)

func NewBVM()() *bvmCreate.BVM {
	vm, errs := bvmCreate.CreateBVM(bvm, nil, deploys, oxygens, nil)
	if errs != nil {
		log.Println(errs)
		return nil
	}
	return bvm
}
