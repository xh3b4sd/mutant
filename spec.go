package mutant

// Interface defines how implementations of permutation patterns need to be made
// accessible to the user. Below is an example to illustrate the usage of any
// permutation implementation.
//
//     var p mutant.Interface
//     {
//         c := perm.Config{
//             Capacity: []int{1, 1, 1},
//         }
//
//         p, err = perm.New(c)
//         if err != nil {
//             ...
//         }
//     }
//
//     for {
//         select {
//         case <-p.Check():
//             ...
//         default:
//             for _, a := range as {
//                 l := p.Index()
//                 ...
//             }
//
//             p.Shift()
//         }
//     }
//
type Interface interface {
	Check() <-chan struct{}
	Index() []int
	Reset()
	Shift()
}
