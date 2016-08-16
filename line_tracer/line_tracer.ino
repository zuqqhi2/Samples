#define PBS1 2
#define PWM9 9
#define PWM10 10
#define analogPin0 0
#define analogPin1 1
#define analogPin2 2

int value0, value1, value2;

void setup()
{
  pinMode(PBS1,INPUT);
  pinMode(PWM9,OUTPUT);
  pinMode(PWM10,OUTPUT);

  Serial.begin(9600);
}

void loop()
{
  digitalWrite(PWM9, LOW);
  digitalWrite(PWM10,LOW);
  int s1 = digitalRead(PBS1);
  Serial.print("s1: ");
  Serial.println(s1);
  if (s1 == LOW)
  {
    while(1)
    {
      value0 = analogRead(analogPin0);
      value1 = analogRead(analogPin1);
      value2 = analogRead(analogPin2);
      Serial.print("value0, value1, value2: ");
      Serial.println(value0);
      Serial.println(value1);
      Serial.println(value2);
      if (value1 < 100 && value2 < 100)
      {
        analogWrite(PWM9, value0 / 4);
        analogWrite(PWM10, value0 / 4);
      }
      if (value1 > 200 && value2 < 100)
      {
        analogWrite(PWM9, value0 / 16);
        analogWrite(PWM10, value0 / 4);
      }
      if (value1 < 100 && value2 > 200)
      {
        analogWrite(PWM9, value0 / 4);
        analogWrite(PWM10, value0 / 16);
      }
    }
  }
}
