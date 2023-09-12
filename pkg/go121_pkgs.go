//go:build go1.21
// +build go1.21

package pkg

import (
	_ "github.com/goplus/igop/pkg/archive/tar"
	_ "github.com/goplus/igop/pkg/archive/zip"
	_ "github.com/goplus/igop/pkg/bufio"
	_ "github.com/goplus/igop/pkg/bytes"
	_ "github.com/goplus/igop/pkg/compress/bzip2"
	_ "github.com/goplus/igop/pkg/compress/flate"
	_ "github.com/goplus/igop/pkg/compress/gzip"
	_ "github.com/goplus/igop/pkg/compress/lzw"
	_ "github.com/goplus/igop/pkg/compress/zlib"
	_ "github.com/goplus/igop/pkg/container/heap"
	_ "github.com/goplus/igop/pkg/container/list"
	_ "github.com/goplus/igop/pkg/container/ring"
	_ "github.com/goplus/igop/pkg/context"
	_ "github.com/goplus/igop/pkg/crypto"
	_ "github.com/goplus/igop/pkg/crypto/aes"
	_ "github.com/goplus/igop/pkg/crypto/cipher"
	_ "github.com/goplus/igop/pkg/crypto/des"
	_ "github.com/goplus/igop/pkg/crypto/dsa"
	_ "github.com/goplus/igop/pkg/crypto/ecdh"
	_ "github.com/goplus/igop/pkg/crypto/ecdsa"
	_ "github.com/goplus/igop/pkg/crypto/ed25519"
	_ "github.com/goplus/igop/pkg/crypto/elliptic"
	_ "github.com/goplus/igop/pkg/crypto/hmac"
	_ "github.com/goplus/igop/pkg/crypto/md5"
	_ "github.com/goplus/igop/pkg/crypto/rand"
	_ "github.com/goplus/igop/pkg/crypto/rc4"
	_ "github.com/goplus/igop/pkg/crypto/rsa"
	_ "github.com/goplus/igop/pkg/crypto/sha1"
	_ "github.com/goplus/igop/pkg/crypto/sha256"
	_ "github.com/goplus/igop/pkg/crypto/sha512"
	_ "github.com/goplus/igop/pkg/crypto/subtle"
	_ "github.com/goplus/igop/pkg/crypto/tls"
	_ "github.com/goplus/igop/pkg/crypto/x509"
	_ "github.com/goplus/igop/pkg/crypto/x509/pkix"
	_ "github.com/goplus/igop/pkg/database/sql"
	_ "github.com/goplus/igop/pkg/database/sql/driver"
	_ "github.com/goplus/igop/pkg/debug/buildinfo"
	_ "github.com/goplus/igop/pkg/debug/dwarf"
	_ "github.com/goplus/igop/pkg/debug/elf"
	_ "github.com/goplus/igop/pkg/debug/gosym"
	_ "github.com/goplus/igop/pkg/debug/macho"
	_ "github.com/goplus/igop/pkg/debug/pe"
	_ "github.com/goplus/igop/pkg/debug/plan9obj"
	_ "github.com/goplus/igop/pkg/embed"
	_ "github.com/goplus/igop/pkg/encoding"
	_ "github.com/goplus/igop/pkg/encoding/ascii85"
	_ "github.com/goplus/igop/pkg/encoding/asn1"
	_ "github.com/goplus/igop/pkg/encoding/base32"
	_ "github.com/goplus/igop/pkg/encoding/base64"
	_ "github.com/goplus/igop/pkg/encoding/binary"
	_ "github.com/goplus/igop/pkg/encoding/csv"
	_ "github.com/goplus/igop/pkg/encoding/gob"
	_ "github.com/goplus/igop/pkg/encoding/hex"
	_ "github.com/goplus/igop/pkg/encoding/json"
	_ "github.com/goplus/igop/pkg/encoding/pem"
	_ "github.com/goplus/igop/pkg/encoding/xml"
	_ "github.com/goplus/igop/pkg/errors"
	_ "github.com/goplus/igop/pkg/expvar"
	_ "github.com/goplus/igop/pkg/flag"
	_ "github.com/goplus/igop/pkg/fmt"
	_ "github.com/goplus/igop/pkg/go/ast"
	_ "github.com/goplus/igop/pkg/go/build"
	_ "github.com/goplus/igop/pkg/go/build/constraint"
	_ "github.com/goplus/igop/pkg/go/constant"
	_ "github.com/goplus/igop/pkg/go/doc"
	_ "github.com/goplus/igop/pkg/go/doc/comment"
	_ "github.com/goplus/igop/pkg/go/format"
	_ "github.com/goplus/igop/pkg/go/importer"
	_ "github.com/goplus/igop/pkg/go/parser"
	_ "github.com/goplus/igop/pkg/go/printer"
	_ "github.com/goplus/igop/pkg/go/scanner"
	_ "github.com/goplus/igop/pkg/go/token"
	_ "github.com/goplus/igop/pkg/go/types"
	_ "github.com/goplus/igop/pkg/hash"
	_ "github.com/goplus/igop/pkg/hash/adler32"
	_ "github.com/goplus/igop/pkg/hash/crc32"
	_ "github.com/goplus/igop/pkg/hash/crc64"
	_ "github.com/goplus/igop/pkg/hash/fnv"
	_ "github.com/goplus/igop/pkg/hash/maphash"
	_ "github.com/goplus/igop/pkg/html"
	_ "github.com/goplus/igop/pkg/html/template"
	_ "github.com/goplus/igop/pkg/image"
	_ "github.com/goplus/igop/pkg/image/color"
	_ "github.com/goplus/igop/pkg/image/color/palette"
	_ "github.com/goplus/igop/pkg/image/draw"
	_ "github.com/goplus/igop/pkg/image/gif"
	_ "github.com/goplus/igop/pkg/image/jpeg"
	_ "github.com/goplus/igop/pkg/image/png"
	_ "github.com/goplus/igop/pkg/index/suffixarray"
	_ "github.com/goplus/igop/pkg/io"
	_ "github.com/goplus/igop/pkg/io/fs"
	_ "github.com/goplus/igop/pkg/io/ioutil"
	_ "github.com/goplus/igop/pkg/log"
	_ "github.com/goplus/igop/pkg/log/slog"
	_ "github.com/goplus/igop/pkg/maps"
	_ "github.com/goplus/igop/pkg/math"
	_ "github.com/goplus/igop/pkg/math/big"
	_ "github.com/goplus/igop/pkg/math/bits"
	_ "github.com/goplus/igop/pkg/math/cmplx"
	_ "github.com/goplus/igop/pkg/math/rand"
	_ "github.com/goplus/igop/pkg/mime"
	_ "github.com/goplus/igop/pkg/mime/multipart"
	_ "github.com/goplus/igop/pkg/mime/quotedprintable"
	_ "github.com/goplus/igop/pkg/net"
	_ "github.com/goplus/igop/pkg/net/http"
	_ "github.com/goplus/igop/pkg/net/http/cgi"
	_ "github.com/goplus/igop/pkg/net/http/cookiejar"
	_ "github.com/goplus/igop/pkg/net/http/fcgi"
	_ "github.com/goplus/igop/pkg/net/http/httptest"
	_ "github.com/goplus/igop/pkg/net/http/httptrace"
	_ "github.com/goplus/igop/pkg/net/http/httputil"
	_ "github.com/goplus/igop/pkg/net/http/pprof"
	_ "github.com/goplus/igop/pkg/net/mail"
	_ "github.com/goplus/igop/pkg/net/netip"
	_ "github.com/goplus/igop/pkg/net/smtp"
	_ "github.com/goplus/igop/pkg/net/textproto"
	_ "github.com/goplus/igop/pkg/net/url"
	_ "github.com/goplus/igop/pkg/os"
	_ "github.com/goplus/igop/pkg/os/exec"
	_ "github.com/goplus/igop/pkg/os/signal"
	_ "github.com/goplus/igop/pkg/os/user"
	_ "github.com/goplus/igop/pkg/path"
	_ "github.com/goplus/igop/pkg/path/filepath"
	_ "github.com/goplus/igop/pkg/plugin"
	_ "github.com/goplus/igop/pkg/reflect"
	_ "github.com/goplus/igop/pkg/regexp"
	_ "github.com/goplus/igop/pkg/regexp/syntax"
	_ "github.com/goplus/igop/pkg/runtime"
	_ "github.com/goplus/igop/pkg/runtime/coverage"
	_ "github.com/goplus/igop/pkg/runtime/debug"
	_ "github.com/goplus/igop/pkg/runtime/metrics"
	_ "github.com/goplus/igop/pkg/runtime/pprof"
	_ "github.com/goplus/igop/pkg/runtime/trace"
	_ "github.com/goplus/igop/pkg/slices"
	_ "github.com/goplus/igop/pkg/sort"
	_ "github.com/goplus/igop/pkg/strconv"
	_ "github.com/goplus/igop/pkg/strings"
	_ "github.com/goplus/igop/pkg/sync"
	_ "github.com/goplus/igop/pkg/sync/atomic"
	_ "github.com/goplus/igop/pkg/syscall"
	_ "github.com/goplus/igop/pkg/testing"
	_ "github.com/goplus/igop/pkg/text/scanner"
	_ "github.com/goplus/igop/pkg/text/tabwriter"
	_ "github.com/goplus/igop/pkg/text/template"
	_ "github.com/goplus/igop/pkg/text/template/parse"
	_ "github.com/goplus/igop/pkg/time"
	_ "github.com/goplus/igop/pkg/time/tzdata"
	_ "github.com/goplus/igop/pkg/unicode"
	_ "github.com/goplus/igop/pkg/unicode/utf16"
	_ "github.com/goplus/igop/pkg/unicode/utf8"
)
