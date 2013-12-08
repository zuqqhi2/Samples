function [s,fs,bits]=wavread(wavefile)

fid=fopen(wavefile, 'rb', 'ieee-le');

% 'RIFF' chunk
RiffChunkID=fread(fid, 4, 'uchar');
RiffChunkSize=fread(fid, 1, 'uint32');
RiffFormType=fread(fid, 4, 'uchar');

% 'fmt' chunk
fmtChunkID=fread(fid, 4, 'uchar');
fmtChunkSize=fread(fid, 1, 'uint32');
fmtWaveFormatType=fread(fid, 1, 'uint16');
fmtChannel=fread(fid, 1, 'uint16');
fmtSamplesPerSec=fread(fid, 1, 'uint32');
fmtBytesPerSec=fread(fid, 1, 'uint32');
fmtBlockSize=fread(fid, 1, 'uint16');
fmtBitsPerSample=fread(fid, 1, 'uint16')

% 'data' chunk
dataChunkID=fread(fid, 4, 'uchar');
dataChunkSize=fread(fid, 1, 'uint32');
if fmtWaveFormatType==1
	if fmtChannel==1
		if fmtBitsPerSample==8
			length_of_s=dataChunkSize;
			s=fread(fid,length_of_s, 'uchar');
			s=(s-128)/128;
		elseif fmtBitsPerSample==16
			length_of_s=dataChunkSize/2;
			s=fread(fid, length_of_s, 'int16');
			s=s/32768;
		end
	end
end
fclose(fid);
fs=fmtSamplesPerSec;
bits=fmtBitsPerSample;
