package lctree

import "testing"

func TestDeserialize(t *testing.T) {
	if tree := Deserialize("[]"); tree != nil {
		t.Fatalf("in: [], got: %v", tree)
	}

	ts := []string{
		"[1,2,3]",
		"[1,null,2,3]",
		"[5,4,7,3,null,2,null,-1,null,9]",
		"[5,1,4,null,null,3,6]",
	}

	for _, tc := range ts {
		if tree := Deserialize(tc); tree == nil {
			t.Fatalf("in: %v, got: %v", tc, tree)
		}
	}
}

func TestSerialize(t *testing.T) {
	tsEmpty := []string{
		"[]",
	}

	for _, tc := range tsEmpty {
		in, want := tc, tc
		if got := Serialize(Deserialize(in)); got != want {
			t.Fatalf("in: %v, want: %v, got: %v", in, want, got)
		}
	}

	ts := []string{
		"[1,2,3]",
		"[1,null,2,3]",
		"[5,4,7,3,null,2,null,-1,null,9]",
		"[5,1,4,null,null,3,6]",
	}

	for _, tc := range ts {
		in, want := tc, tc
		if got := Serialize(Deserialize(in)); got != want {
			t.Fatalf("in: %v, want: %v, got: %v", in, want, got)
		}
	}
}

type TcDOT struct {
	In   string
	Want string
}

func TestDOT(t *testing.T) {
	tsEmpty := []TcDOT{
		{
			In: "[]",
			Want: `digraph {}
`,
		},
	}

	for _, tc := range tsEmpty {
		in, want := tc.In, tc.Want
		tree := Deserialize(in)
		if got := tree.DOT(); got != want {
			t.Fatalf("in: %v\nwant:\n%v,\ngot:\n%v", in, want, got)
		}
	}

	ts := []TcDOT{
		{
			In: "[1,2,3]",
			Want: `digraph {
graph [ordering="out"];
1;
2;
3;
1 -> 2;
1 -> 3;
}
`,
		},
		{
			In: "[1,null,2,3]",
			Want: `digraph {
graph [ordering="out"];
1;
1.0 [label="", width=.1, style=invis];
2;
2.1 [label="", width=.1, style=invis];
3;
1 -> 1.0 [style=invis];
1 -> 2;
2 -> 3;
2 -> 2.1 [style=invis];
}
`,
		},
		{
			In: "[5,4,7,3,null,2,null,-1,null,9]",
			Want: `digraph {
graph [ordering="out"];
5;
4;
4.1 [label="", width=.1, style=invis];
7;
7.1 [label="", width=.1, style=invis];
3;
3.1 [label="", width=.1, style=invis];
2;
2.1 [label="", width=.1, style=invis];
-1;
9;
5 -> 4;
5 -> 7;
4 -> 3;
4 -> 4.1 [style=invis];
7 -> 2;
7 -> 7.1 [style=invis];
3 -> -1;
3 -> 3.1 [style=invis];
2 -> 9;
2 -> 2.1 [style=invis];
}
`,
		},
		// {
		//  In: "[5,1,4,null,null,3,6]",
		// },
	}

	for _, tc := range ts {
		in, want := tc.In, tc.Want
		tree := Deserialize(in)
		if got := tree.DOT(); got != want {
			t.Fatalf("in: %v\nwant:\n%v,\ngot:\n%v", in, want, got)
		}
	}

}
