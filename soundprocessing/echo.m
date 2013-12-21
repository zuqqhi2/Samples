clear;
[s0,fs,bits]=wavread('source/thermo.wav');
length_of_s=length(s0);
ap=0.25;
as=0.4;
dp=round(0.100*fs);
ds=round(0.050*fs);
rs=10;
a=zeros(1,rs+2);
a(1)=1;
a(2)=ap;
for k=1:rs,
	a(k+2)=a(k+1)*as;
end
d=zeros(1,rs+2);
d(1)=0;
d(2)=dp;
for k=1:rs,
	d(k+2)=d(k+1)+ds;
end
s1=zeros(1,length_of_s);
for n=1:length_of_s,
	for k=1:rs+2,
		if n-d(k) > 0
			s1(n)=s1(n)+a(k)*s0(n-d(k));
		end
	end
end
wavwrite(s1,fs,bits,'dest/echo.wav');
