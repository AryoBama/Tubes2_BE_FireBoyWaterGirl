package types

type RecipeNode struct {
	Element     string       `json:"element"`    
	Combination []IngredientPair     `json:"combination,omitempty"`
	Children    []RecipeNode  `json:"children,omitempty"`
}

type RecipeTree struct {
	Root      string     `json:"root"`
	Recipe    RecipeNode  `json:"recipe"`
}

type Combo struct{
	ID     int      `json:"id"`
    Inputs []string `json:"inputs"`
    Output string   `json:"output"`
}



func NewRecipeNode(element string, combination []IngredientPair, children []RecipeNode) *RecipeNode{

	return &RecipeNode{
		Element : element,
		Combination: combination,
		Children: children,
	}

}

func (n *RecipeNode) AddChild(child *RecipeNode){

	n.Children = append(n.Children, *child)

}
