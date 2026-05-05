#!/usr/bin/env bash
set -euo pipefail

BASE="http://localhost:8080"
PASS=0
FAIL=0

check() {
  local method="$1" path="$2" expect="$3"
  status=$(curl -s -o /dev/null -w "%{http_code}" -X "$method" "$BASE$path")
  if [ "$status" = "$expect" ]; then
    echo "  PASS $method $path → $status"
    PASS=$((PASS + 1))
  else
    echo "  FAIL $method $path → $status (expected $expect)"
    FAIL=$((FAIL + 1))
  fi
}

check_json() {
  local method="$1" path="$2" body="$3" expect="$4"
  status=$(curl -s -o /dev/null -w "%{http_code}" -X "$method" -H "Content-Type: application/json" -d "$body" "$BASE$path")
  if [ "$status" = "$expect" ]; then
    echo "  PASS $method $path → $status"
    PASS=$((PASS + 1))
  else
    echo "  FAIL $method $path → $status (expected $expect)"
    FAIL=$((FAIL + 1))
  fi
}

echo "=== PulseOS E2E Smoke Test ==="

check GET /healthz 200

# User
echo "--- User ---"
check_json POST /api/v1/users/onboarding '{"name":"测试用户","age":28,"gender":"male","height_cm":175,"weight_kg":70,"primary_goal":"maintain"}' 201
check GET /api/v1/users/profile 200
check GET /api/v1/users/settings 200
check GET /api/v1/users/stats 200

# Diet
echo "--- Diet ---"
check GET /api/v1/diet/plan/today 200
check_json POST /api/v1/diet/photo-upload?filename=meal.jpg '{}' 201
check_json POST /api/v1/diet/analyze '{"image_url":"/test.jpg","meal_type":"lunch","manual_items":["鸡胸肉沙拉"]}' 200
check GET /api/v1/diet/records 200
check_json POST /api/v1/diet/records '{"image_url":"/test.jpg","meal_type":"lunch","manual_items":["糙米饭"]}' 201

# Activity
echo "--- Activity ---"
check_json POST /api/v1/activity/records '{"activity_type":"walking","minutes":30,"intensity":"moderate","steps":3000}' 201
check GET /api/v1/activity/records 200
check GET /api/v1/activity/today 200
check GET /api/v1/activity/week 200

# Meditation
echo "--- Meditation ---"
check_json POST /api/v1/meditation/sessions '{"mode_key":"calm","duration_s":300,"audio_key":"rain"}' 201
check GET /api/v1/meditation/today 200

# Sleep
echo "--- Sleep ---"
check_json POST /api/v1/sleep/sessions/start '{"audio_url":"/audio/sleep.wav"}' 201
check_json POST /api/v1/sleep/sessions/end '{"session_id":"sleep-1"}' 200
check GET /api/v1/sleep/today 200

# Home / Scoring
echo "--- Home ---"
check GET /api/v1/home/dashboard 200

echo ""
echo "=== Results: $PASS passed, $FAIL failed ==="
[ "$FAIL" -eq 0 ]
