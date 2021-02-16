RSpec.describe 'not to method' do 
  it 'checks for the inverse of a matcher' do 
    expect(5).not_to eq(10)
    expect([1, 2, 3]).not_to equal([1, 2, 3])
    expect(5).not_to be_even
  end
end
