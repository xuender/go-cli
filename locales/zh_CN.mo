Þ    ?        Y         p  J   q  -   ¼  &   ê  f     #   x  K     I   è  &   2     Y  \   m      Ê  -   ë          *     A     _  (   l          ¦      ¹     Ú     ì     	  ]   !	     	  !   	     ±	  B   ¿	  A   
  ,   D
     q
     
     
  E   ª
  )   ð
  *        E     Q     k               ¯  \   ½          /     =     N     ]     p            	   £     ­     º     Ç     Ð     ì               *     8  (   I  .  r  S   ¡  +   õ  "   !  d   D  "   ©  S   Ì  R      1   s     ¥  X   ¾  !     -   9     g     t  !        §  '   ´     Ü     é     ö               0  U   @          £     Â  6   Ï  <     (   C     l            5   ¼  #   ò  $        ;     H  *   ^               ²  Q   Â          '     5     F     W     h     y               ¬     À     Ó  #   á            *   #     N     e  $   v     #          
      8              5   9   &           $                   6       ?         >   0         +          3   %   /                             2   ;   =   .       )   7                                  <         "      :   *   1               	       -                   (   !           '                 ,   4      # Create cmd
  go-cli g c cmd
  # Create cobra
  go-cli g c cmd -t cobra   # Create example
  go-cli g e pkg/source.go   # Create interface
  go-cli g i Book   # Create message
  go-cli g p pb/Book
  # Create enum
  go-cli g p BookType -t enum -o pb/book.proto   # Create struct
  go-cli g s Book   # Create test
  go-cli g t pkg/source.go
  # Create path
  go-cli g t pkg   # Init project
  go-cli init
  # Init github config
 go-cli init github   # Watch dir
  go-cli watch [command] CLI tool for Golang CLI tool for Golang

  Generate structures, tests, examples, initialize projects, etc.

		%s Convert struct to other structs. Create a new struct function by other struct. Create templates Create templates to %s Debug mode, display debug log Generate cmd Generate command support cobra and flag. Generate example Generate interface Generate interface and comments. Generate protobuf Generate protobuf and comments. Generate source code Generate source code.

Including commands, tests, examples, struct, interface, protobuf, etc. Generate struct Generate struct and new function. Generate test Generate test examples for exposed functions in file or directory. Generating unit tests for exposed functions in file or directory. Golang source file %s struct num is not one. Init Golang project Init gitee config Init github config Initialize the Golang project and create default configuration files. Initialize the gitee configuration files. Initialize the github configuration files. Output file Select type cobra or flag Select type message or enum Struct related Struct related commands. Watch and Run Watch to the directory and run a command. If the directory is modified, restart the command. create %s templates. create %s: %s create Proto: %s create cmd: %s create function %s create test: %s create: %s %s dir is not exist: %s exist: %s init dir: %s init file %s init: %s license: APACHE2, BSD3, MIT no duplicate name field not found: %s select language: en, zh set debug: %v set language: %s watch path, default is current directory Project-Id-Version: 1.1.12
Report-Msgid-Bugs-To: 
PO-Revision-Date: 2023-07-13 15:16+0800
Last-Translator: ender <xuender@139.com>
Language-Team: ender <xuender@139.com>
Language: zh_CN
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
X-Generator: Poedit 3.0.1
   # çæå½ä»¤
  go-cli g c cmd
  # çæ cobra å½ä»¤
  go-cli g c cmd -t cobra   # çæä¾å­
  go-cli g e pkg/source.go   # çææ¥å£
  go-cli g i Book   # çææ¶æ¯
  go-cli g p pb/Book
  # create enum
  go-cli g p BookType -t enum -o pb/book.proto   # çæç»æ
  go-cli g s Book   # çææµè¯
  go-cli g t pkg/source.go
  # æ ¹æ®ç®å½çæ
  go-cli g t pkg   # åå§åé¡¹ç®
  go-cli init
  # åå§å github éç½®
  go-cli init github   # çå¬ç®å½å¹¶è¿è¡
  go-cli watch [å½ä»¤] Go è¯­è¨å½ä»¤è¡å·¥å· Go è¯­è¨å½ä»¤è¡å·¥å·

  çæç»æãæµè¯ãä¾å­ï¼åå§åé¡¹ç®ç­ã

		%s å°ç»æè½¬æ¢ä¸ºå¶ä»ç»æã éè¿å¶ä»ç»æåå»ºä¸ä¸ªæ°çç»æã åå»ºæ¨¡æ¿ åå»ºæ¨¡æ¿: %s è°è¯æ¨¡å¼ï¼æ¾ç¤ºè°è¯æ¥å¿ çæå½ä»¤ çæå½ä»¤ï¼æ¯æ cobra å flagã çæä¾å­ çææ¥å£ çææ¥å£åæ³¨éã çæ protobuf çæ protobuf. çææºä»£ç  çææºä»£ç ã

åæ¬å½ä»¤ãæµè¯ãä¾å­ãç»æãæ¥å£ãprotobufç­ã çæç»æ çæç»æåæ°å»ºå½æ°ã çææµè¯ çææä»¶æç®å½ä¸­å¬å¼å½æ°çæµè¯æä¾å­ æ ¹æ®ç®å½ææä»¶ä¸­çå¬å¼å½æ°çæååæµè¯ã Golangæºç¨åº %s ç»æå·ä¸æ¯ä¸ã åå§åGoè¯­è¨é¡¹ç® åå§å gitee.com éç½® åå§å  github.com éç½® åå§ågoè¯­è¨é¡¹ç®å¹¶åå»ºé»è®¤éç½®æä»¶ã åå§å gitee.com éç½®æä»¶ã åå§å github.com éç½®æä»¶ã è¾åºæä»¶ éæ© cobra æ flag éæ©ç±»åæ¶æ¯(message)ææä¸¾(enum) ç»æç¸å³ ä¸ç»æç¸å³çå½ä»¤ã çå¬å¹¶è¿è¡ çå¬ç®å½å¹¶è¿è¡å½ä»¤ãå¦æç®å½å·²ä¿®æ¹ï¼è¯·éæ°å¯å¨è¯¥å½ä»¤ã åå»º: %s æ¨¡æ¿. åå»º %s: %s åå»ºç»æ: %s åå»ºå½ä»¤: %s åå»ºå½æ°: %s åå»ºæµè¯: %s åå»º: %s %s ç®å½ä¸å­å¨: %s å·²ç»å­å¨: %s åå§åç®å½: %s åå§åæä»¶ %s åå§å: %s è®¸å¯åè®®ï¼APACHE2ãBSD3ãMIT æ éå¤å­æ®µ æªæ¾å°: %s éæ©è¯­è¨: è±è¯­(en_US), ä¸­æ(zh_CN) è®¾ç½®è°è¯æ¨¡å¼: %v è®¾ç½®è¯­è¨: %s çè§è·¯å¾ï¼é»è®¤ä¸ºå½åç®å½ 