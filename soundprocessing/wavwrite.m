function wavwrite(s,fs,bits,wavefile)

fid=fopen(wavefile, 'wb', 'ieee-le');
length_of_s=length(s);
if bits==8
	dataChunkSize=length_of_s;
elseif bits==16
	dataChunkSize=length_of_s*2;
end

RiffChunkID='RIFF';
RiffChunkSize=dataChunkSize+36;
RiffFormType='WAVE';
fmtChunkID='fmt ';
fmtChunkSize=16;
fmtWaveFormatType=1;
fmtChannel=1;
fmtSamplesPerSec=fs;
fmtBytesPerSec=fs*bits/8;
fmtBlockSize=bits/8;
fmtBitsPerSample=bits;
dataChunkID='data';

% 'RIFF' chunk
fwrite(fid, RiffChunkID, 'uchar');
fwrite(fid, RiffChunkSize, 'uint32');
fwrite(fid, RiffFormType, 'uchar');

% 'fmt' chunk
fwrite(fid, fmtChunkID, 'uchar');
fwrite(fid, fmtChunkSize, 'uint32');
fwrite(fid, fmtWaveFormatType, 'uint16');
fwrite(fid, fmtChannel, 'uint16');
fwrite(fid, fmtSamplesPerSec, 'uint32');
fwrite(fid, fmtBytesPerSec, 'uint32');
fwrite(fid, fmtBlockSize, 'uint16');
fwrite(fid, fmtBitsPerSample, 'uint16');

% 'data' chunk
fwrite(fid, dataChunkID, 'uchar');
fwrite(fid, dataChunkSize, 'uint32');
if bits==8
	s=round((s+1)*255/2);
	fwrite(fid,s,'uchar');
elseif bits==16
	s=round((s+1)*65535/2)-32768;
	fwrite(fid,s,'int16');
end
fclose(fid);
