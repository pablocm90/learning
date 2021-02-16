RSpec.describe 'satisfy matcher' do 
  subject { 'racecar' }

  it 'is a palindrome' do 
    expect(subject).to satisfy('be a palindrome') { |value| value == value.reverse }
  end
end
