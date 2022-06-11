FROM public.ecr.aws/lambda/go:1.x

RUN yum install -y epel-release chromium

CMD ["bin/film_search"]