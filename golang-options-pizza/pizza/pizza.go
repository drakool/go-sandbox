package Pizza

type Pizza struct {
    Dough    string
    Sauce    string
    Cheese   string
    Toppings []string
}

type PizzaOptions struct {
    Dough    string
    Sauce    string
    Cheese   string
    Toppings []string
}

type PizzaOption func(*PizzaOptions)

func WithDough(dough string) PizzaOption {
    return func(po *PizzaOptions) {
        po.Dough = dough
    }
}

func WithSauce(sauce string) PizzaOption {
    return func(po *PizzaOptions) {
        po.Sauce = sauce
    }
}

func WithCheese(cheese string) PizzaOption {
    return func(po *PizzaOptions) {
        po.Cheese = cheese
    }
}

func WithToppings(toppings []string) PizzaOption {
    return func(po *PizzaOptions) {
        po.Toppings = toppings
    }
}

func NewPizza(options ...PizzaOption) *Pizza {
    opts := &PizzaOptions{}
    for _, option := range options {
        option(opts)
    }

    pizza := &Pizza{
        Dough:    opts.Dough,
        Sauce:    opts.Sauce,
        Cheese:   opts.Cheese,
        Toppings: opts.Toppings,
    }

    return pizza
}
