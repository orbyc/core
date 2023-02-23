publish:
	@npm publish

add:
	@git add .

commit:
	@git commit -m "release ${RELEASE_VERSION}"

push:
	@git push origin master

createtag:
	@git tag ${RELEASE_VERSION}

pushtag:
	@git push origin ${RELEASE_VERSION}

release: publish add commit push createtag pushtag