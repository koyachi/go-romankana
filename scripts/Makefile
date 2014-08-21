all: clone_romankana_rb k2r.go r2k.go

clone_romankana_rb:
	git clone https://github.com/ymrl/romankana.git tmp/romankana

k2r.go:
	@echo 'package romankana\n' > tmp/$@
	@echo 'k2r := map[string]string{' >> tmp/$@
	@cat tmp/romankana/lib/romankana/k2r.rb | awk '/=>/{ gsub(/\047/,"\""); print "\t"$$1": " $$3 }' >> tmp/$@
	@echo '}' >> tmp/$@

# これだけじゃだけめだ。最初にflattenするか。
r2k.go:
	@echo 'package romankana\n' > tmp/$@
	@echo 'r2k := map[string]string{' >> tmp/$@
	@cat tmp/romankana/lib/romankana/r2k.rb | awk '/=>/{ print "\t"$$1": " $$3 }' >> tmp/$@
	@echo '}' >> tmp/$@
