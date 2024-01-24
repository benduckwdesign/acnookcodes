package acnookcodes

// Debugging
var DEBUG bool = false
var DEBUG_CPtrShift bool = false

// Code Types
const (
	Famicom int = iota
	Popular
	CardE
	Magazine
	User
	CardEMini
)

// Hit Rates
const (
	HITRATE_EIGHTY_80 = iota
	HITRATE_SIXTY_60
	HITRATE_THIRTY_30
	HITRATE_ZERO_0
	HITRATE_ONE_HUNDRED_100
)

const (
	CARDE_HITRATE_EIGHTY_80 = iota
	CARDE_HITRATE_SIXTY_60
	CARDE_HITRATE_FORTY_40
	CARDE_HITRATE_TWENTY_20
)

// ASCII Codes
const (
	CHAR_INVERT_EXCLAMATION   = iota // ¡
	CHAR_INVERT_QUESTIONMARK         // ¿
	CHAR_DIAERESIS_A                 // Ä
	CHAR_GRAVE_A                     // À
	CHAR_ACUTE_A                     // Á
	CHAR_CIRCUMFLEX_A                // Â
	CHAR_TILDE_A                     // Ã
	CHAR_ANGSTROM_A                  // Å
	CHAR_CEDILLA                     // Ç
	CHAR_GRAVE_E                     // È
	CHAR_ACUTE_E                     // É
	CHAR_CIRCUMFLEX_E                // Ê
	CHAR_DIARESIS_E                  // Ë
	CHAR_GRAVE_I                     // Ì
	CHAR_ACUTE_I                     // Í
	CHAR_CIRCUMFLEX_I                // Î
	CHAR_DIARESIS_I                  // Ï
	CHAR_ETH                         // Ð
	CHAR_TILDE_N                     // Ñ
	CHAR_GRAVE_O                     // Ò
	CHAR_ACUTE_O                     // Ó
	CHAR_CIRCUMFLEX_O                // Ô
	CHAR_TILDE_O                     // Õ
	CHAR_DIARESIS_O                  // Ö
	CHAR_OE                          // Œ
	CHAR_GRAVE_U                     // Ù
	CHAR_ACUTE_U                     // Ú
	CHAR_CIRCUMFLEX_U                // Û
	CHAR_DIARESIS_U                  // Ü
	CHAR_LOWER_BETA                  // β
	CHAR_THORN                       // Þ
	CHAR_GRAVE_a                     // à
	CHAR_SPACE                       // ' '
	CHAR_EXCLAMATION                 // !
	CHAR_QUOTATION                   // "
	CHAR_ACUTE_a                     // á
	CHAR_CIRCUMFLEX_a                // â
	CHAR_PERCENT                     // %
	CHAR_AMPERSAND                   // &
	CHAR_APOSTROPHE                  // '
	CHAR_OPEN_PARENTHESIS            // (
	CHAR_CLOSE_PARENTHESIS           // )
	CHAR_TILDE                       // ~
	CHAR_SYMBOL_HEART                // ♥
	CHAR_COMMA                       // ,
	CHAR_DASH                        // -
	CHAR_PERIOD                      // .
	CHAR_SYMBOL_MUSIC_NOTE           // ♪
	CHAR_ZERO                        // 0
	CHAR_ONE                         // 1
	CHAR_TWO                         // 2
	CHAR_THREE                       // 3
	CHAR_FOUR                        // 4
	CHAR_FIVE                        // 5
	CHAR_SIX                         // 6
	CHAR_SEVEN                       // 7
	CHAR_EIGHT                       // 8
	CHAR_NINE                        // 9
	CHAR_COLON                       // :
	CHAR_SYMBOL_DROPLET              // ☂
	CHAR_LESS_THAN                   // <
	CHAR_EQUALS                      // =
	CHAR_GREATER_THAN                // >
	CHAR_QUESTIONMARK                // ?
	CHAR_AT_SIGN                     // @
	CHAR_A                           // A
	CHAR_B                           // B
	CHAR_C                           // C
	CHAR_D                           // D
	CHAR_E                           // E
	CHAR_F                           // F
	CHAR_G                           // G
	CHAR_H                           // H
	CHAR_I                           // I
	CHAR_J                           // J
	CHAR_K                           // K
	CHAR_L                           // L
	CHAR_M                           // M
	CHAR_N                           // N
	CHAR_O                           // O
	CHAR_P                           // P
	CHAR_Q                           // Q
	CHAR_R                           // R
	CHAR_S                           // S
	CHAR_T                           // T
	CHAR_U                           // U
	CHAR_V                           // V
	CHAR_W                           // W
	CHAR_X                           // X
	CHAR_Y                           // Y
	CHAR_Z                           // Z
	CHAR_TILDE_a                     // ã
	CHAR_SYMBOL_ANNOYED              // ☹
	CHAR_DIARESIS_a                  // ä
	CHAR_ANGSTROM_a                  // å
	CHAR_UNDERSCORE                  // _
	CHAR_LOWER_CEDILLA               // ç
	CHAR_a                           // a
	CHAR_b                           // b
	CHAR_c                           // c
	CHAR_d                           // d
	CHAR_e                           // e
	CHAR_f                           // f
	CHAR_g                           // g
	CHAR_h                           // h
	CHAR_i                           // i
	CHAR_j                           // j
	CHAR_k                           // k
	CHAR_l                           // l
	CHAR_m                           // m
	CHAR_n                           // n
	CHAR_o                           // o
	CHAR_p                           // p
	CHAR_q                           // q
	CHAR_r                           // r
	CHAR_s                           // s
	CHAR_t                           // t
	CHAR_u                           // u
	CHAR_v                           // v
	CHAR_w                           // w
	CHAR_x                           // x
	CHAR_y                           // y
	CHAR_z                           // z
	CHAR_GRAVE_e                     // è
	CHAR_ACUTE_e                     // é
	CHAR_CIRCUMFLEX_e                // ê
	CHAR_DIARESIS_e                  // ë
	CHAR_CONTROL_CODE                // 
	CHAR_MESSAGE_TAG                 // 
	CHAR_GRAVE_i                     // ì
	CHAR_ACUTE_i                     // í
	CHAR_CIRCUMFLEX_i                // î
	CHAR_DIARESIS_i                  // ï
	CHAR_INTERPUNCT                  // ·
	CHAR_LOWER_ETH                   // ð
	CHAR_TILDE_n                     // ñ
	CHAR_GRAVE_o                     // ò
	CHAR_ACUTE_o                     // ó
	CHAR_CIRCUMFLEX_o                // ô
	CHAR_TILDE_o                     // õ
	CHAR_DIARESIS_o                  // ö
	CHAR_oe                          // ø
	CHAR_GRAVE_u                     // ù
	CHAR_ACUTE_u                     // ú
	CHAR_HYPHEN                      // -
	CHAR_CIRCUMFLEX_u                // û
	CHAR_DIARESIS_u                  // ü
	CHAR_ACUTE_y                     // ý
	CHAR_DIARESIS_y                  // ÿ
	CHAR_LOWER_THORN                 // þ
	CHAR_ACUTE_Y                     // ý
	CHAR_BROKEN_BAR                  // ¦
	CHAR_SILCROW                     // §
	CHAR_FEMININE_ORDINAL            // ª
	CHAR_MASCULINE_ORDINAL           // º
	CHAR_DOUBLE_VERTICAL_BAR         // ¦
	CHAR_LATIN_MU                    // µ
	CHAR_SUPERSCRIPT_THREE           // ³
	CHAR_SUPERSCRIPT_TWO             // ²
	CHAR_SUPRESCRIPT_ONE             // ¹
	CHAR_MACRON_SYMBOL               // ¯
	CHAR_LOGICAL_NEGATION            // ¡
	CHAR_ASH                         // £
	CHAR_LOWER_ASH                   // ¥
	CHAR_INVERT_QUOTATION            // ¢
	CHAR_GUILLEMET_OPEN              // «
	CHAR_GUILLEMET_CLOSE             // »
	CHAR_SYMBOL_SUN                  // ¿
	CHAR_SYMBOL_CLOUD                // É
	CHAR_SYMBOL_UMBRELLA             // Û
	CHAR_SYMBOL_WIND                 // î
	CHAR_SYMBOL_SNOWMAN              // è
	CHAR_LINES_CONVERGE_RIGHT        // ï
	CHAR_LINES_CONVERGE_LEFT         // ò
	CHAR_FORWARD_SLASH               // ô
	CHAR_INFINITY                    // ¥
	CHAR_CIRCLE                      // û
	CHAR_CROSS                       // ¢
	CHAR_SQUARE                      // ª
	CHAR_TRIANGLE                    // ò
	CHAR_PLUS                        // ¦
	CHAR_SYMBOL_LIGTNING             // ¦
	CHAR_MARS_SYMBOL                 // ¥
	CHAR_VENUS_SYMBOL                // É
	CHAR_SYMBOL_FLOWER               // Û
	CHAR_SYMBOL_STAR                 // î
	CHAR_SYMBOL_SKULL                // è
	CHAR_SYMBOL_SURPRISE             // ï
	CHAR_SYMBOL_HAPPY                // ò
	CHAR_SYMBOL_SAD                  // ô
	CHAR_SYMBOL_ANGRY                // ¥
	CHAR_SYMBOL_SMILE                // ¢
	CHAR_DIMENSION_SIGN              // ¡
	CHAR_OBELUS_SIGN                 // £
	CHAR_SYMBOL_HAMMER               // ¥
	CHAR_SYMBOL_RIBBON               // £
	CHAR_SYMBOL_MAIL                 // ¥
	CHAR_SYMBOL_MONEY                // £
	CHAR_SYMBOL_PAW                  // ¥
	CHAR_SYMBOL_SQUIRREL             // £
	CHAR_SYMBOL_CAT                  // ¥
	CHAR_SYMBOL_RABBIT               // £
	CHAR_SYMBOL_OCTOPUS              // ¥
	CHAR_SYMBOL_COW                  // £
	CHAR_SYMBOL_PIG                  // ¥
	CHAR_NEW_LINE                    // £
	CHAR_SYMBOL_FISH                 // ¥
	CHAR_SYMBOL_BUG                  // £
	CHAR_SEMICOLON                   // ;
	CHAR_HASHTAG                     // #
	CHAR_SPACE_2                     // Short space
	CHAR_SPACE_3                     // Wide space
	CHAR_SYMBOL_KEY                  // ☁
	// Begin EU-only symbols, unused in AC
	CHAR_LEFT_QUOTATION   // «
	CHAR_RIGHT_QUOTATION  // »
	CHAR_LEFT_APOSTROPHE  // ‘
	CHAR_RIGHT_APOSTROPHE // ’
	CHAR_ETHEL            // Ð
	CHAR_LOWER_ETHEL      // ð
	CHAR_ORDINAL_e        // ª
	CHAR_ORDINAL_er       // º
	CHAR_ORDINAL_re       // ¿
	CHAR_BACKSLASH        // /
	// Unused characters
	CHAR_223 // Unused
	CHAR_224 // Unused
	CHAR_225 // Unused
	CHAR_226 // Unused
	CHAR_227 // Unused
	CHAR_228 // Unused
	CHAR_229 // Unused
	CHAR_230 // Unused
	CHAR_231 // Unused
	CHAR_232 // Unused
	CHAR_233 // Unused
	CHAR_234 // Unused
	CHAR_235 // Unused
	CHAR_236 // Unused
	CHAR_237 // Unused
	CHAR_238 // Unused
	CHAR_239 // Unused
	CHAR_240 // Unused
	CHAR_241 // Unused
	CHAR_242 // Unused
	CHAR_243 // Unused
	CHAR_244 // Unused
	CHAR_245 // Unused
	CHAR_246 // Unused
	CHAR_247 // Unused
	CHAR_248 // Unused
	CHAR_249 // Unused
	CHAR_250 // Unused
	CHAR_251 // Unused
	CHAR_252 // Unused
	CHAR_253 // Unused
	CHAR_254 // Unused
	CHAR_255 // Unused
)

const (
	PARAM_STRING_SIZE         = 8
	PASSWORD_CHAR_BITS        = 6
	PASSWORD_DATA_BITS        = 8
	PASSWORD_DATA_SIZE        = 21
	PASSWORD_BITMIXKEY_IDX    = 1
	PASSWORD_RSA_KEY01_IDX    = 15
	PASSWORD_RSA_EXPONENT_IDX = 5
	PASSWORD_RSA_BITSAVE_IDX  = 20
)

var PASSWORD_STR_SIZE = (PASSWORD_DATA_SIZE * PASSWORD_DATA_BITS) / PASSWORD_CHAR_BITS
