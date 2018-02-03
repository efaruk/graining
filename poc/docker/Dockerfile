# This docker file is only to verify it can run without any runtime dependency
FROM scratch
COPY assets/bin/* /
# Second copy line is necessary otherwise it is coping every thing in /var/lock
COPY assets/empty /var/lock/ 

CMD ["/autoc"]
