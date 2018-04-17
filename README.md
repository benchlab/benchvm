# BenchVM

A distributed virtual machine layer within the BenchChain MultiChain. BenchVM, also known as BVM is relied upon by the BenchChain for executing smart contracts and is built into BenchChain's Core Node and the GoLang implementation (BenGo)

We built BVM with both EOS and Ethereum's EVM in mind. We built BVM to follow the specifications and functionalities offered by both EOS and Ethereum's underlying smart contract protocols. Any EOS or Ethereum-based smart contracts can be deployed on any Bench Multichain.

We also set out to minimize the bytecode size due to on-chain storage. All DAPPs on the Ethereum network in their smallest EVMState are stored at 32 total bytes.

As an example, a boolean within a EVM-based DAPP is stored on-chain as :
`0x0000000000000001 or 0x0000000000000000`

## Our Approach To Distributed VMs
BVM-based DAPPs run on a prefix-based specification, where the data above would in turn look like this:
`0x0101`

The total specification includes a `size byte` of `1`, a `value byte` of `1` and the `parameterStore` as one byte as well. To figure these numbers, you would need a size byte of length 1 but if not, the size byte can always be altered by the first byte of the BVM Smart Contract.


## Reduction Formula
For a size parameter of length n
`downsize := 32 - bytesLength - s`

## Note
This library is not for individual use. It is used within BenchChain Core and BenGo implementation.

## Roadmap
- Mapping needs to be more recursive
- Integrating other smart contract platforms to work with the BVM Spec.
