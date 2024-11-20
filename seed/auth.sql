INSERT INTO two_factor_types (name, description)
SELECT 'Google Authenticator', 'Time-based One-Time Password (via Google Authenticator)'
WHERE NOT EXISTS (SELECT 1 FROM two_factor_types WHERE name = 'Google Authenticator');

INSERT INTO two_factor_types (name, description)
SELECT 'SMS', 'Send OTP via SMS'
WHERE NOT EXISTS (SELECT 1 FROM two_factor_types WHERE name = 'SMS');

INSERT INTO two_factor_types (name, description)
SELECT 'Email', 'Send OTP via Email'
WHERE NOT EXISTS (SELECT 1 FROM two_factor_types WHERE name = 'Email');

INSERT INTO two_factor_types (name, description)
SELECT 'Push', 'Push notification to mobile app'
WHERE NOT EXISTS (SELECT 1 FROM two_factor_types WHERE name = 'Push');

INSERT INTO two_factor_types (name, description)
SELECT 'Backup code', 'Use one of 10 backup codes. After usage backup code becomes unavailable'
WHERE NOT EXISTS (SELECT 1 FROM two_factor_types WHERE name = 'Backup code');