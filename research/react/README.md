# React とは

## 仮想 DOM

開発者は、具体的な DOM 操作を行わなくても OK

### 素の Javascript

User - Javascript - DOM

### React

User - **React** - Javascript - DOM

- React はデータの更新を検知して、変更後の DOM 操作をエミュレートする。
  変更前の状態も残しておくことで、差分だけ DOM 操作処理を行う。

- 仮想 DOM は Javascript レベルの賢いコーディング方法

## データと UI の同期

### 大抵の Javascript

イベント - データの取得 - UI 更新

DOM 更新とそれ以外の処理が一体となりがち

### React

イベント - データの取得 - データの更新 - **UI 更新(React)**

- ライブラリ(React)がデータと UI の橋渡しを実装
- UI が自動的に切り替わる

## コンポーネント指向
