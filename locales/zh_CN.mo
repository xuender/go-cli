Þ    4      ¼  G   \      x  J   y  -   Ä  &   ò  f     #     K   ¤  I   ð     :  V   N     ¥     ¶     Í     ë  (   ø     !     2      E     f     x       ]   ­       !        =  B   K  A        Ð     ä     ö  E   		  )   O	  *   y	     ¤	     °	     Ê	     æ	     û	     	
     
     )
     9
     G
  	   \
     f
     s
     
     
     ¥
     ³
     Ë
     Ù
    ê
  S     +   Z  "     d   ©  "     S   1  R        Ø  R   ñ     D     Q  !   b       '        ¹     Æ     Ó     ì     ü       U        s            6   ¬  <   ã           8     S  5   p  #   ¦  $   Ê     ï     ü  *        =     P     ^     o                    ³     Ä     Ø     ë  #   ù       *   +     V     m        '                    1      $                /   %                   	               &              ,   -      0   4   +   3             #   (   
                                 *       .      "                      )            2       !             # Create cmd
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
 go-cli init github CLI tool for Golang CLI tool for Golang

  Generate structures, tests, examples, initialize projects, etc. Create templates Create templates to %s Debug mode, display debug log Generate cmd Generate command support cobra and flag. Generate example Generate interface Generate interface and comments. Generate protobuf Generate protobuf and comments. Generate source code Generate source code.

Including commands, tests, examples, struct, interface, protobuf, etc. Generate struct Generate struct and new function. Generate test Generate test examples for exposed functions in file or directory. Generating unit tests for exposed functions in file or directory. Init Golang project Init gitee config Init github config Initialize the Golang project and create default configuration files. Initialize the gitee configuration files. Initialize the github configuration files. Output file Select type cobra or flag Select type message or enum create %s templates. create %s: %s create Proto: %s create cmd: %s create test: %s create: %s %s dir is not exist: %s exist: %s init dir: %s init file %s init: %s license: APACHE2, BSD3, MIT not found: %s select language: en, zh set debug: %v set language: %s Project-Id-Version: 1.2.11
Report-Msgid-Bugs-To: 
PO-Revision-Date: 2023-04-03 11:27+0800
Last-Translator: Automatically generated
Language-Team: none
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
  go-cli init github Go è¯­è¨å½ä»¤è¡å·¥å· Go è¯­è¨å½ä»¤è¡å·¥å·

  çæç»æãæµè¯ãä¾å­ï¼åå§åé¡¹ç®ç­ã åå»ºæ¨¡æ¿ åå»ºæ¨¡æ¿: %s è°è¯æ¨¡å¼ï¼æ¾ç¤ºè°è¯æ¥å¿ çæå½ä»¤ çæå½ä»¤ï¼æ¯æ cobra å flagã çæä¾å­ çææ¥å£ çææ¥å£åæ³¨éã çæ protobuf çæ protobuf. çææºä»£ç  çææºä»£ç ã

åæ¬å½ä»¤ãæµè¯ãä¾å­ãç»æãæ¥å£ãprotobufç­ã çæç»æ çæç»æåæ°å»ºå½æ°ã çææµè¯ çææä»¶æç®å½ä¸­å¬å¼å½æ°çæµè¯æä¾å­ æ ¹æ®ç®å½ææä»¶ä¸­çå¬å¼å½æ°çæååæµè¯ã åå§åGoè¯­è¨é¡¹ç® åå§å gitee.com éç½® åå§å  github.com éç½® åå§ågoè¯­è¨é¡¹ç®å¹¶åå»ºé»è®¤éç½®æä»¶ã åå§å gitee.com éç½®æä»¶ã åå§å github.com éç½®æä»¶ã è¾åºæä»¶ éæ© cobra æ flag éæ©ç±»åæ¶æ¯(message)ææä¸¾(enum) åå»º: %s æ¨¡æ¿. åå»º %s: %s åå»ºç»æ: %s åå»ºå½ä»¤: %s åå»ºæµè¯: %s åå»º: %s %s ç®å½ä¸å­å¨: %s å·²ç»å­å¨: %s åå§åç®å½: %s åå§åæä»¶ %s åå§å: %s è®¸å¯åè®®ï¼APACHE2ãBSD3ãMIT æªæ¾å°: %s éæ©è¯­è¨: è±è¯­(en_US), ä¸­æ(zh_CN) è®¾ç½®è°è¯æ¨¡å¼: %v è®¾ç½®è¯­è¨: %s 