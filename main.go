package main

func main() {
	bc := new_block_chain()
	bc.add_block("sample1")
	bc.add_block("sample2")
	bc.add_block("sample3")
	bc.add_block("sample3")
	bc.add_block("sample3")
	bc.add_block("sample3")
	bc.mine()
	bc.mine()
	// bc.mine()
	// bc.mine()
	// bc.mine()
	// bc.mine()
	// bc.mine()
	bc.display_chain()
}
