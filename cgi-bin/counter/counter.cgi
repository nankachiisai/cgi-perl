#!/bin/perl

use strict;
use warnings;

# ファイルを開いてカウント値を取得
open(my $fh_in, '<', 'counter/count.txt') or die "File opening was failed: $!";
my $count = <$fh_in>;
chomp($count);
close($fh_in);

# インクリメント
$count++;

# ファイルを開いてカウント値を書き戻す
open(my $fh_out, '>', 'counter/count.txt') or die "File opening was failed: $!";
print $fh_out "$count";

# HTML出力
print "Content-type: text/html\n\n";
print "<html lang=\"ja\"><head><meta charset=\"utf-8\"><title>カウンター</title></head><body>";
print "<p>";
for (my $i = 7; $i >= 0; $i--) {
    my $num = 0;
    if ($i == 0) {
        $num = $count % 10;
    } else {
        $num = ($count / (10 ** $i)) % 10;
    }
    print "<img src=\"/img/number_$num.png\">";
}
print "</p>";
print "</body></html>";
