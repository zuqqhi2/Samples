clear;
window_size=16384;
[s0,fs,bits]=wavread('source/thermo.wav');
s1=zeros(1,window_size);
for n=1:window_size,
	s1(n)=s0(n);
end
wavwrite(s1,fs,bits,'dest/thermo.wav');
