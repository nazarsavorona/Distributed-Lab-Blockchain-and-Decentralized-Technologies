package ec_point

import (
	"math/big"
	"reflect"
	"testing"
)

func TestAddECPoints(t *testing.T) {
	p := new(ECPoint)
	n := new(big.Int)
	m := new(big.Int)
	n, _ = n.SetString("5955751385846415552691510251172147524855841759259746329344872112054088090368567558489174224710360475979610786089000070229367361974620147989668756088331814429", 10)
	m, _ = m.SetString("3607062370101963573933943048150495311826377203515354188559873974163406182234489407026269066910260080025972849251988568707375791772468913697371892599991412842", 10)

	type args struct {
		a ECPoint
		b ECPoint
	}
	tests := []struct {
		name string
		args args
		want ECPoint
	}{
		{
			name: "Double",
			args: args{
				a: p.BasePointGGet(),
				b: p.BasePointGGet(),
			},
			want: p.DoubleECPoints(p.BasePointGGet()),
		},
		{
			name: "Using multiplication",
			args: args{
				a: p.ScalarMult(p.BasePointGGet(), *big.NewInt(3)),
				b: p.ScalarMult(p.BasePointGGet(), *big.NewInt(4)),
			},
			want: p.ScalarMult(p.BasePointGGet(), *big.NewInt(7)),
		},
		{
			name: "Expected from value",
			args: args{
				a: p.ScalarMult(p.BasePointGGet(), *big.NewInt(7)),
				b: p.ScalarMult(p.BasePointGGet(), *big.NewInt(11)),
			},
			want: p.ECPointGen(n, m),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.AddECPoints(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddECPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasePointGGet(t *testing.T) {
	p := new(ECPoint)
	n := new(big.Int)
	m := new(big.Int)
	n, _ = n.SetString("2661740802050217063228768716723360960729859168756973147706671368418802944996427808491545080627771902352094241225065558662157113545570916814161637315895999846", 10)
	m, _ = m.SetString("3757180025770020463545507224491183603594455134769762486694567779615544477440556316691234405012945539562144444537289428522585666729196580810124344277578376784", 10)

	tests := []struct {
		name string
		want ECPoint
	}{
		{
			name: "Generate base point",
			want: p.ECPointGen(n, m),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.BasePointGGet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasePointGGet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleECPoints(t *testing.T) {
	p := new(ECPoint)
	n := new(big.Int)
	m := new(big.Int)
	n, _ = n.SetString("901472452850866198617673658578940391618730359691416279093035377195377079020397774511960179466499271590922803070095487687963115616363390991670183687363590205", 10)
	m, _ = m.SetString("3281327921582527507824747162491172657218985358085640380741461489720525905953211486053138004786012424348623853685340634287932228687534583594738661002099038978", 10)

	type args struct {
		a ECPoint
	}
	tests := []struct {
		name string
		args args
		want ECPoint
	}{
		{
			name: "Double base point",
			args: args{
				a: p.BasePointGGet(),
			},
			want: p.ECPointGen(n, m),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.DoubleECPoints(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoubleECPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestECPointToString(t *testing.T) {
	p := new(ECPoint)
	n := new(big.Int)
	m := new(big.Int)
	n, _ = n.SetString("123", 10)
	m, _ = m.SetString("345", 10)

	type args struct {
		point ECPoint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String representation",
			args: args{
				point: p.ECPointGen(n, m),
			},
			want: "(123; 345)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := p.ECPointToString(tt.args.point); gotS != tt.want {
				t.Errorf("ECPointToString() = %v, want %v", gotS, tt.want)
			}
		})
	}
}

func TestIsOnCurveCheck(t *testing.T) {
	p := new(ECPoint)
	n := new(big.Int)
	m := new(big.Int)
	n, _ = n.SetString("901472452850866198617673658578940391618730359691416279093035377195377079020397774511960179466499271590922803070095487687963115616363390991670183687363590205", 10)
	m, _ = m.SetString("3281327921582527507824747162491172657218985358085640380741461489720525905953211486053138004786012424348623853685340634287932228687534583594738661002099038978", 10)

	type args struct {
		a ECPoint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "On curve",
			args: args{
				a: p.ECPointGen(n, m),
			},
			want: true,
		},
		{
			name: "Not on curve",
			args: args{
				a: p.ECPointGen(big.NewInt(123), big.NewInt(345)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.IsOnCurveCheck(tt.args.a); got != tt.want {
				t.Errorf("IsOnCurveCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScalarMult(t *testing.T) {
	p := new(ECPoint)
	k := big.NewInt(5)
	n := new([4]*big.Int)

	for i := range n {
		n[i] = big.NewInt(0)
	}

	n[0], _ = n[0].SetString("2661740802050217063228768716723360960729859168756973147706671368418802944996427808491545080627771902352094241225065558662157113545570916814161637315895999846", 10)
	n[1], _ = n[1].SetString("3757180025770020463545507224491183603594455134769762486694567779615544477440556316691234405012945539562144444537289428522585666729196580810124344277578376784", 10)
	n[2], _ = n[2].SetString("1356490565846790255739168589204180496138457916455904646417394236994934604955446470133745046215752716497315667741727281784411650125266382302052918851202396280", 10)
	n[3], _ = n[3].SetString("4664604347668020762970169188647728988711329258092833895899217703661992685928809983190581180899656432001123782446415831494072942427109258745206503403498279115", 10)

	type args struct {
		a ECPoint
		k big.Int
	}
	tests := []struct {
		name string
		args args
		want ECPoint
	}{
		{
			name: "Multiply by 5",
			args: args{
				a: p.ECPointGen(n[0], n[1]),
				k: *k,
			},
			want: p.ECPointGen(n[2], n[3]),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.ScalarMult(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScalarMult() = %v, want %v", got, tt.want)
			}
		})
	}
}
