# SGS IMAGE

SGS IMAGE はシナリオスクリプトの DSL の実験リポジトリです。

scenario モジュールは外部です。対して app モジュールはゲームエンジン側です。
scenario 側の独立性を確保します。scenario 側は同一で app 側のゲームエンジン
を切り替えられるようにします。

app 側から scenario を制御しますが scenario の中には立ち入りません。
app と scenario とを繋ぐのはトランジションです。

トランジションです。トランジションはたとえば、カメラの移動を表現します。
app はトランジションを使ってアニメーションします。トランジションは
終わった後の値を示します。そうすることで途中からの再生をしやすくするほか、