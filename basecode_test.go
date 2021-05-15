package basecode

import (
	"strconv"
	"testing"
	"time"
)

func TestEncoding_Encode(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "hello world Encode",
			args: args{src: "hello world"},
			want: "GCCAGCGGGCUAGCUAGCUUACAAGUGUGCUUGUACGCUAGCGA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StdEncoding.Encode(tt.args.src); got != tt.want {
				t.Errorf("Encoding.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncoding_Decode(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "hello world Decode",
			args: args{src: "GCCAGCGGGCUAGCUAGCUUACAAGUGUGCUUGUACGCUAGCGA"},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StdEncoding.Decode(tt.args.src); got != tt.want {
				t.Errorf("Encoding.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOtherCodebookEncoding_Encode(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "hello world Encode",
			args: args{src: "hello world"},
			want: "夏秋秋春夏秋夏夏夏秋东春夏秋东春夏秋东东春秋春春夏东夏东夏秋东东夏东春秋夏秋东春夏秋夏春",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncoding([]string{"春", "夏", "秋", "东"})
			if got := e.Encode(tt.args.src); got != tt.want {
				t.Errorf("Encoding.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOtherCodebookEncoding_Decode(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "hello world Encode",
			args: args{src: "夏秋秋春夏秋夏夏夏秋东春夏秋东春夏秋东东春秋春春夏东夏东夏秋东东夏东春秋夏秋东春夏秋夏春"},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NewEncoding([]string{"春", "夏", "秋", "东"})
			if got := e.Decode(tt.args.src); got != tt.want {
				t.Errorf("Encoding.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkStdEncoding(t *testing.B) {
	t.StartTimer()

	t.Run("stressTest", func(b *testing.B) {
		for i := 0; i < t.N; i++ {
			StdEncoding.Encode(strconv.Itoa(i))
		}
	})

	t.StopTimer()

}

func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			StdEncoding.Encode(strconv.Itoa(time.Now().Second()))
		}
	})
}
